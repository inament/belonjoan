package main

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bagus2x/sumux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "crud.db")
	stmt, _ := db.Prepare(
		`CREATE TABLE IF NOT EXISTS belonjoan (
		item_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		price NUMERIC NOT NULL
		)`,
	)

	stmt.Exec()

	r := sumux.NewMux()
	r.Log = false

	InitBelonjoanHandler(db, r)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		dir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(dir, "client"))
		fs := http.StripPrefix("/", http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
