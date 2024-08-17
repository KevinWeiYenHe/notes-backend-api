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
# Current Goals

- Setup server
- Setup routing
- Setup handlers

- Create database migrations for notes





---
Special thanks to Alex Edwards to his book, Let's Go and Let's Go Further