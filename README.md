# Notes - Backend API

A backend API for a notes app. A future frontend for the application will be located [here](https://github.com/KevuTheDev/notes-frontend).

---
## Summary
This is a JSON API using Go as a backend language. 

Uses the following Go library to setup the backend:
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) // **Router**


Uses [PostgreSQL](https://www.postgresql.org/) as the **database** of choice

Uses Docker to run PostgreSQL

---
# Endpoints

| Method | URL Pattern | Action |
| -- | -- | -- |
| **GET** | /v1/ping | Ping route to test if server is active | 
| **GET** | /v1/healthcheck | Show application health and version information | 
| **GET** | /v1/notes | Show the details of all notes | 
| **POST** | /v1/notes | Create a new note |
| **GET** | /v1/notes/:id | Show the details of a specific note | 
| **PATCH** | /v1/notes/:id | Update the details of a specific note | 
| **DELETE** | /v1/notes/:id | Delete a specific note | 

---
# Notes Model
```GO
type Note struct {
	ID           int64     `json:"id"`                // unique id for the note
	CreatedAt    time.Time `json:"created_at"`        // when the note was created
	LastUpdateAt time.Time `json:"last_updated_at"`   // when the note was last updated
	Title        string    `json:"title"`             // title of note
	Content      string    `json:"content,omitempty"` // content of note
	Tags         []string  `json:"tags,omitempty"`    // tags of note
	Version      int32     `json:"version"`           // number of times the note was updated
}
```

# Notes Form
Title:   string   - Cannot be empty
Content: string   - Can be empty
Tags:    []string - Can be empty

CreatedAt:    time.Time (Assigned at POST)
LastUpdateAt: time.Time (Assigned at POST, updated at PATCH)

---
# Current Goals
- Create database migrations for notes





---
Special thanks to Alex Edwards to his book, Let's Go and Let's Go Further