package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/KevuTheDev/notes-backend-api/internal/validator"
	"github.com/lib/pq"
)

// Must be in LOWERCASE
var ValidTags = map[string]string{
	"yes":   "",
	"no":    "",
	"world": "",
}

type Note struct {
	ID           int64     `json:"id"`                // unique id for the note
	CreatedAt    time.Time `json:"created_at"`        // when the note was created
	LastUpdateAt time.Time `json:"last_updated_at"`   // when the note was last updated
	Title        string    `json:"title"`             // title of note
	Content      string    `json:"content,omitempty"` // content of note
	Tags         []string  `json:"tags,omitempty"`    // tags of note
	Version      int32     `json:"version"`           // number of times the note was updated
}

func ValidateNote(v *validator.Validator, note *Note) {
	v.Check(note.Title != "", "title", "must be provided")
	v.Check(len(note.Title) <= 500, "title", "must not be more than 500 bytes long")

	// TODO
	// VALID TAGS ARE OFF
	// v.Check(validator.PermittedValues(note.Tags, ValidTags), "tags", "invalid tags accepted")

	v.Check(validator.Unique(note.Tags), "tags", "must not contain duplicate values")
}

// Define a NoteModel struct type which wraps a sql.DB connection pool
type NoteModel struct {
	DB *sql.DB
}

func (n NoteModel) Insert(note *Note) error {
	stmt := `
		INSERT INTO notes (title, content, tags)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, last_updated_at, version`

	args := []any{note.Title, note.Content, pq.Array(note.Tags)}

	return n.DB.QueryRow(stmt, args...).Scan(&note.ID, &note.CreatedAt, &note.LastUpdateAt, &note.Version)
}

func (n NoteModel) Get(id int64) (*Note, error) {
	stmt := `
		SELECT id, created_at, last_updated_at, title, content, tags, version 
		FROM notes
		WHERE id = $1`

	var note Note

	err := n.DB.QueryRow(stmt, id).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.LastUpdateAt,
		&note.Title,
		&note.Content,
		pq.Array(&note.Tags),
		&note.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &note, nil
}

func (n NoteModel) Update(note *Note) error {
	stmt := `
		UPDATE notes
		SET title = $1, content = $2, tags = $3, last_updated_at = NOW(), version = version + 1
		WHERE id = $4
		RETURNING version, last_updated_at`

	args := []any{
		note.Title,
		note.Content,
		pq.Array(note.Tags),
		note.ID,
	}

	return n.DB.QueryRow(stmt, args...).Scan(&note.Version, &note.LastUpdateAt)
}

func (n NoteModel) Delete(id int64) error {
	return nil
}
