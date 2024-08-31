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

# Database
```SQL
-- Database Creation
CREATE DATABASE notebook;

-- Create role notebook 
CREATE ROLE notebook WITH LOGIN PASSWORD 'pa55word'; -- Remember to change this unsecure password 

--- Add citext extension
CREATE EXTENSION IF NOT EXISTS citext;

--- Grant privilage to user on specific database
GRANT ALL PRIVILEGES ON DATABASE EXAMPLE_DB TO EXAMPLE_USER;
--- Switch to database
\c EXAMPLE_DB postgres
--- Grant privilage on schema to user
GRANT ALL ON SCHEMA public TO EXAMPLE_USER;


GRANT ALL PRIVILEGES ON DATABASE notebook TO notebook;

GRANT ALL ON SCHEMA public TO notebook;

```

---
# Current Goals
- Create database migrations for notes



---
# .env FILE
```ini
# POSTGRES DSN string
NOTEBOOK_DB_DSN="postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable"
# SMTP Server information
SMTP_HOST="sandbox.smtp.mailtrap.io"
SMTP_PORT=25
SMTP_USERNAME=""
SMTP_PASSWORD=""
SMTP_SENDER=""
```

---
# Database migration tool
[migrate](https://github.com/golang-migrate/migrate)

```bash
# to migrate a database up/down/goto X
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

# create a migration
migrate create -seq -ext='.sql' -dir='./migrations' migration_process


```

```sql
--- copy onto the psql console
\c notebook postgres
\COPY notes (title, content, tags) FROM  './db/note_data.csv' DELIMITER '|';
```

---
Special thanks to Alex Edwards to his book, Let's Go and Let's Go Further