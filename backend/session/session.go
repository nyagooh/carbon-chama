package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

// Session represents a user session
type Session struct {
	ID        string
	UserID    int
	Username  string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// Store manages sessions
type Store struct {
	sessions map[string]*Session
	mutex    sync.RWMutex
}

// NewStore creates a new session store
func NewStore() *Store {
	return &Store{
		sessions: make(map[string]*Session),
	}
}

// Generate creates a new session for a user
func (s *Store) Generate(userID int, username string) (*Session, error) {
	// Generate random session ID
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	id := base64.URLEncoding.EncodeToString(b)

	// Create session with 24 hour expiry
	session := &Session{
		ID:        id,
		UserID:    userID,
		Username:  username,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// Store session
	s.mutex.Lock()
	s.sessions[id] = session
	s.mutex.Unlock()

	return session, nil
}

// Get retrieves a session by ID
func (s *Store) Get(id string) (*Session, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	session, exists := s.sessions[id]
	if !exists {
		return nil, false
	}
	
	// Check if session has expired
	if time.Now().After(session.ExpiresAt) {
		delete(s.sessions, id)
		return nil, false
	}
	
	return session, true
}

// Delete removes a session
func (s *Store) Delete(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.sessions, id)
}

// SetCookie adds a session cookie to the response
func SetCookie(w http.ResponseWriter, session *Session) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Expires:  session.ExpiresAt,
	}
	http.SetCookie(w, cookie)
}

// GetSessionFromRequest retrieves the session from the request
func GetSessionFromRequest(r *http.Request, store *Store) (*Session, bool) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, false
	}
	
	return store.Get(cookie.Value)
}

// ClearCookie removes the session cookie
func ClearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
