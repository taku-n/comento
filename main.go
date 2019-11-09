package main

import "database/sql"
import "encoding/json"
import . "fmt"
import "net/http"

import _ "github.com/mattn/go-sqlite3"

type Data struct {
	Name string `json:"name"`
	Comment string `json:"comment"`
}

func get_all(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := sql.Open("sqlite3", "comento.db")
	if err != nil {
		Fprintln(w, "Failed to open the database file.")
	}

	rows, err := db.Query(`SELECT name, comment FROM Comento;`)
	if err != nil {
		Fprintln(w, "Failed to get data.")
	}
	defer rows.Close()

	var data = []Data{}
	for rows.Next() {
		var name string
		var comment string

		err := rows.Scan(&name, &comment)
		if err != nil {
			Fprintln(w, "Failed to scan data.")
		}

		data = append(data, Data{name, comment})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		Fprintln(w, "Failed to marshal data: ", err)
	}
	Fprint(w, string(jsonData))
}

func get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := sql.Open("sqlite3", "comento.db")
	if err != nil {
		Fprintln(w, "Failed to open the database file.")
	}

	rows, err := db.Query(`SELECT name, comment FROM Comento WHERE thread = ?;`,
			r.Form["thread"][0])
	if err != nil {
		Fprintln(w, "Failed to get data: ", err)
	}
	defer rows.Close()

	var data = []Data{}
	for rows.Next() {
		var name string
		var comment string

		err := rows.Scan(&name, &comment)
		if err != nil {
			Fprintln(w, "Failed to scan data.")
		}

		data = append(data, Data{name, comment})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		Fprintln(w, "Failed to marshal data: ", err)
	}
	Fprint(w, string(jsonData))
}

func post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Fprintln(w, r.Form)
	Fprintln(w, r.Form["thread"][0])

	db, err := sql.Open("sqlite3", "comento.db")
	if err != nil {
		Fprintln(w, "Failed to open the database file.")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Comento
			(id      INTEGER,
			 thread  TEXT,
			 name    TEXT,
			 comment TEXT,
		 	 PRIMARY KEY (id));`,
	)
	if err != nil {
		Fprintln(w, "Failed to create a table.")
	}

	_, err = db.Exec(`INSERT INTO Comento (thread, name, comment) VALUES (?, ?, ?)`,
		r.Form["thread"][0], r.Form["name"][0], r.Form["comment"][0])
	if err != nil {
		Fprintln(w, "Failed to insert a record.")
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/get_all", get_all)
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)

	err := server.ListenAndServe()
	if err != nil {
		Println("ListenAndServe() error: ", err)
	}
}
