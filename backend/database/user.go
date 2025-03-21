package database

import (
	"errors"
	"sync"
)

// User represents a user in the system
type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

// UserDB is a simple in-memory user database
type UserDB struct {
	users     map[string]*User // Email to user mapping
	usersByID map[int]*User    // ID to user mapping
	nextID    int
	mutex     sync.RWMutex
}

// NewUserDB creates a new user database with some dummy users
func NewUserDB() *UserDB {
	db := &UserDB{
		users:     make(map[string]*User),
		usersByID: make(map[int]*User),
		nextID:    1,
	}

	// Add some dummy users
	db.AddUser("admin", "admin@example.com", "admin123")
	db.AddUser("user1", "user1@example.com", "password123")
	db.AddUser("user2", "user2@example.com", "password123")

	return db
}

// AddUser adds a new user to the database
func (db *UserDB) AddUser(username, email, password string) (*User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Check if user already exists
	if _, exists := db.users[email]; exists {
		return nil, errors.New("user with this email already exists")
	}

	// Create new user
	user := &User{
		ID:       db.nextID,
		Username: username,
		Email:    email,
		Password: password, // In a real app, this would be hashed
	}

	// Store user
	db.users[email] = user
	db.usersByID[user.ID] = user
	db.nextID++

	return user, nil
}

// GetUserByEmail retrieves a user by email
func (db *UserDB) GetUserByEmail(email string) (*User, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.users[email]
	return user, exists
}

// GetUserByID retrieves a user by ID
func (db *UserDB) GetUserByID(id int) (*User, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.usersByID[id]
	return user, exists
}

// Authenticate checks if the email and password are valid
func (db *UserDB) Authenticate(email, password string) (*User, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.users[email]
	if !exists || user.Password != password {
		return nil, false
	}

	return user, true
}

// UpdateUser updates a user's information
func (db *UserDB) UpdateUser(id int, username, email, password string) (*User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, exists := db.usersByID[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	// If email is changing, check if new email is already in use
	if email != user.Email {
		if _, exists := db.users[email]; exists {
			return nil, errors.New("email already in use")
		}
		delete(db.users, user.Email)
		db.users[email] = user
	}

	// Update user
	user.Username = username
	user.Email = email
	if password != "" {
		user.Password = password
	}

	return user, nil
}

// DeleteUser removes a user from the database
func (db *UserDB) DeleteUser(id int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, exists := db.usersByID[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(db.users, user.Email)
	delete(db.usersByID, id)

	return nil
}
