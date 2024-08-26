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

func (m NoteModel) Insert(note *Note) error {
	stmt := `
		INSERT INTO notes (title, content, tags)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, last_updated_at, version`

	// set info into the args to return back to the client
	args := []any{note.Title, note.Content, pq.Array(note.Tags)}

	// the scan part matches with the RETURNING columns order
	// .Scan
	return m.DB.QueryRow(stmt, args...).Scan(&note.ID, &note.CreatedAt, &note.LastUpdateAt, &note.Version)
}

func (m NoteModel) Get(id int64) (*Note, error) {
	stmt := `
		SELECT id, created_at, last_updated_at, title, content, tags, version 
		FROM notes
		WHERE id = $1`

	// var to information to send back to client
	var note Note

	// columns should match how the statement order
	err := m.DB.QueryRow(stmt, id).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.LastUpdateAt,
		&note.Title,
		&note.Content,
		pq.Array(&note.Tags),
		&note.Version,
	)

	// if error encounter on the process
	if err != nil {
		switch {
		// no rows foudn
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &note, nil
}

func (m NoteModel) Update(note *Note) error {
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

	return m.DB.QueryRow(stmt, args...).Scan(&note.Version, &note.LastUpdateAt)
}

func (m NoteModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM notes
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m NoteModel) Latest() ([]*Note, error) {
	stmt := `
		SELECT id, created_at, last_updated_at, title, content, tags, version 
		FROM notes
		ORDER BY last_updated_at DESC 
		LIMIT 25`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []*Note{}

	for rows.Next() {
		n := &Note{}

		err = rows.Scan(&n.ID, &n.CreatedAt, &n.LastUpdateAt, &n.Title, &n.Content, pq.Array(&n.Tags), &n.Version)
		if err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
