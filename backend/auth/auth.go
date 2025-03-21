package auth

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "142302828852-pk4b7kvu5ame1eglj2oo0hgi2dg2d4cc.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

// Session store
var store = sessions.NewCookieStore([]byte("super-secret-key"))

// userDB is a mock database for demonstration purposes
var userDB = &mockUserDB{}

type mockUserDB struct{}

func (m *mockUserDB) AddUser(username, email, password string) (*User, error) {
	// Mock implementation, replace with actual database logic
	return &User{Username: username, Email: email}, nil
}

func (m *mockUserDB) Authenticate(email, password string) (*User, bool) {
	// Mock implementation, replace with actual database logic
	return &User{Username: "username", Email: email}, true
}

type User struct {
	Username string
	Email    string
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	Url := googleOauthConfig.AuthCodeURL("random-state-token")
	http.Redirect(w, r, Url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "athuration code not found", http.StatusBadRequest)
		return
	}
	token, err := googleOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "failed to exchange token", http.StatusBadRequest)
		return
	}
	// fetch user info from google
	client := googleOauthConfig.Client(r.Context(), token)
	resp, err := client.Get(("https://www.googleapis.com/oauth2/v2/userinfo"))
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var user struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
		return
	}
	// create session
	session, _ := store.Get(r, "sessionname")
	session.Values["Email"] = user.Email
	session.Values["Name"] = user.Name
	session.Values["ProfilePic"] = user.Picture
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	// redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sessionname")
	email, emailExists := session.Values["Email"].(string)
	name, nameExists := session.Values["Name"].(string)
	profilePic, _ := session.Values["ProfilePic"].(string)

	if !emailExists || !nameExists {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("./frontend/templates/dashboard.html"))

	data := struct {
		Email      string
		Username   string
		ProfilePic string
	}{
		Email:      email,
		Username:   name,
		ProfilePic: profilePic,
	}

	tmpl.Execute(w, data)
}

// implement logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sessionname")
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// HandleSignup processes the signup form submission
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate input (basic validation)
	if username == "" || email == "" || password == "" {
		http.Redirect(w, r, "/signup?error=missing_fields", http.StatusSeeOther)
		return
	}

	// Create user
	_, err = userDB.AddUser(username, email, password)
	if err != nil {
		http.Redirect(w, r, "/signup?error=user_exists", http.StatusSeeOther)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// HandleLogin processes the login form submission
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Authenticate user
	user, ok := userDB.Authenticate(email, password)
	if !ok {
		// Authentication failed
		http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
		return
	}

	// Create session
	session, _ := store.Get(r, "sessionname")
	session.Values["Email"] = user.Email
	session.Values["Name"] = user.Username
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	// Redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
