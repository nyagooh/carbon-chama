package serve

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methhod not allowed", http.StatusMethodNotAllowed)
		return
	}
	var page string
	switch r.URL.Path {
	case "/signup":
		page = "signup.html"
	case "/login":
		page = "login.html"
	case "/":
		page = "login.html"
	case "/dashboard":
		page = "dashboard.html"
	default:
		http.NotFound(w, r)
		return
	}

	tmplpath := filepath.Join("frontend", "templates", page)
	tmpl, err := template.ParseFiles(tmplpath)
	if err != nil {
		fmt.Println("Hello Anne You finally found me!!")
		http.Error(w, "internalservererror", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
}
