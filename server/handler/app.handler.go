package handler

import (
	"database/sql"
	"encoding/json"
	"golauth/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	migrate "github.com/rubenv/sql-migrate"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
	router   *mux.Router
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}

	signinHandler := app.SignIn
	signupHandler := app.SignUp
	assignHandler := app.AssignProfile

	signinHandler = disableCors(signinHandler)
	signupHandler = disableCors(signupHandler)
	assignHandler = disableCors(assignHandler)

	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/api/auth/signin", signinHandler)
	router.HandleFunc("/api/auth/signup", signupHandler)
	router.HandleFunc("/api/auth/{user_uuid}/{profile_uuid}", assignHandler)

	app.router = router

	return app
}

func (a *App) Migrate(d *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/mysql",
	}

	n, err := migrate.Exec(d, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Print("Erro ao realizar migrations!\n", err)
	}
	log.Print("Resultado das migrations:", n)
}

func (a *App) Serve() error {
	http.Handle("/api/", a.router)
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
