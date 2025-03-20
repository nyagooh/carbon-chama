package routes

import (
	m "forum/backend/auth"

	s "forum/backend/serve"
	"net/http"
)

// Router initializes and returns a ServeMux with all routes configured
func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Page routes
	mux.HandleFunc("/signup", s.ServePage)
	mux.HandleFunc("/login", s.ServePage)
	mux.HandleFunc("/", s.ServePage)

	// Authentication routes
	mux.HandleFunc("/auth/google", m.HandleGoogleLogin)
	mux.HandleFunc("/auth/callback", m.HandleGoogleCallback)
	mux.HandleFunc("/dashboard", m.HandleDashboard)
	mux.HandleFunc("/logout", m.HandleLogout)

	return mux
}
