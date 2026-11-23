// Package notes keeps notes in memory and serves them over HTTP JSON.
package notes

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ErrNotFound is returned when a note ID does not exist.
var ErrNotFound = errors.New("note not found")

// Note is a single note.
type Note struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// Store keeps notes in memory. It is safe to use from many goroutines,
// because the HTTP server handles every request in its own goroutine.
type Store struct {
	mu     sync.Mutex
	data   map[int]Note
	nextID int
}

// NewStore creates an empty store.
func NewStore() *Store {
	return &Store{data: map[int]Note{}, nextID: 1}
}

// Add saves a new note and returns it with its ID.
func (s *Store) Add(title, body string) Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := Note{ID: s.nextID, Title: title, Body: body, Created: time.Now().UTC()}
	s.data[n.ID] = n
	s.nextID++
	return n
}

// All returns every note, sorted by ID.
func (s *Store) All() []Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Note, 0, len(s.data))
	for _, n := range s.data {
		out = append(out, n)
	}
	slices.SortFunc(out, func(a, b Note) int { return a.ID - b.ID })
	return out
}

// Get returns one note.
func (s *Store) Get(id int) (Note, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	n, ok := s.data[id]
	if !ok {
		return Note{}, ErrNotFound
	}
	return n, nil
}

// Delete removes one note.
func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return ErrNotFound
	}
	delete(s.data, id)
	return nil
}

// Handler returns the HTTP router for this store.
//
//	GET    /notes       every note
//	POST   /notes       {"title": "...", "body": "..."}
//	GET    /notes/{id}  one note
//	DELETE /notes/{id}  delete a note
func Handler(s *Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /notes", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, s.All())
	})
	mux.HandleFunc("POST /notes", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			writeError(w, http.StatusBadRequest, "body must be JSON")
			return
		}
		if strings.TrimSpace(in.Title) == "" {
			writeError(w, http.StatusBadRequest, "title is required")
			return
		}
		writeJSON(w, http.StatusCreated, s.Add(in.Title, in.Body))
	})
	mux.HandleFunc("GET /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "id must be a number")
			return
		}
		n, err := s.Get(id)
		if errors.Is(err, ErrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, n)
	})
	mux.HandleFunc("DELETE /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "id must be a number")
			return
		}
		if err := s.Delete(id); err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
