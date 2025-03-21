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
	mux.HandleFunc("/chama", s.ServePage)
	mux.HandleFunc("/dashboard2", s.ServePage)
	mux.HandleFunc("/save", s.ServePage)

	// Authentication routes
	mux.HandleFunc("/auth/google", m.HandleGoogleLogin)
	mux.HandleFunc("/auth/google/callback", m.HandleGoogleCallback)
	mux.HandleFunc("/auth/login", m.HandleLogin)
	mux.HandleFunc("/auth/signup", m.HandleSignup)
	mux.HandleFunc("/dashboard", m.HandleDashboard)
	mux.HandleFunc("/logout", m.HandleLogout)

	return mux
}
