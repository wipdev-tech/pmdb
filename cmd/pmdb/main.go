// package main is the entry point of the PMDb app
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/google/uuid"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/wipdev-tech/pmdb/internal/database"
)

type service struct {
	db *database.Queries
}

func main() {
	dbURL, dbToken, err := getDBEnv()
	if err != nil {
		log.Fatal(err)
	}

	connURL := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
	db, err := sql.Open("libsql", connURL)
	if err != nil {
		log.Fatal(err)
	}

	s := service{db: database.New(db)}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/add-user", s.handleAddUser)

	fmt.Println("PMDb server let's Go! ")
	if os.Getenv("ENV") == "dev" {
		fmt.Println("Dev server started and running at http://localhost:8080")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}
	fmt.Println("Server started and running")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}

// handleHome is the handler for the home route ("/")
func (s *service) handleHome(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := s.db.ListUsers(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	tmplData := struct {
		Users []database.User
	}{
		Users: dbUsers,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/fragments.html"))
	err = tmpl.Execute(w, tmplData)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *service) handleAddUser(w http.ResponseWriter, r *http.Request) {
	newDisplayName := r.PostFormValue("display-name")
	newUserName := r.PostFormValue("user-name")
	_, err := s.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.NewString(),
		UserName:    newUserName,
		DisplayName: newDisplayName,
	})
	if err != nil {
		log.Fatal(err)
	}

	errMsg := ""
	if err != nil {
		errMsg = "Could not add user :("
	}

	dbUsers, err := s.db.ListUsers(r.Context())
	tmplData := struct {
		Users        []database.User
		ErrorMessage string
	}{
		Users:        dbUsers,
		ErrorMessage: errMsg,
	}

	tmpl := template.Must(template.ParseFiles("templates/users.html"))
	err = tmpl.Execute(w, tmplData)
	if err != nil {
		log.Fatal(err)
	}
}
