package main

import "database/sql"
import "encoding/json"
import . "fmt"
import "net/http"

import _ "github.com/mattn/go-sqlite3"

type DataAll struct {
	Thread  string `json:"thread"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Data struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func root(w http.ResponseWriter, r *http.Request) {
	Println("---- root() begins ----")
	Fprintln(w, "hello, world")
	Println("----  root() ends  ----")
}

func get_all(w http.ResponseWriter, r *http.Request) {
	Println("---- get_all() begins ----")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	Println("---- sql.Open() begins ----")
	db, err := sql.Open("sqlite3", "comento/comento.db")
	if err != nil {
		Println("Failed to open the database file: ", err)
		return
	}

	Println("---- sql.Query() begins ----")
	rows, err := db.Query(`SELECT thread, name, comment FROM Comento;`)
	if err != nil {
		Println("Failed to get data: ", err)
		return
	}
	defer rows.Close()

	Println("---- Scan() begins ----")
	var data = []DataAll{}
	for rows.Next() {
		var thread string
		var name string
		var comment string

		err := rows.Scan(&thread, &name, &comment)
		if err != nil {
			Println("Failed to scan data: ", err)
			return
		}

		data = append(data, DataAll{thread, name, comment})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		Println("Failed to marshal data: ", err)
		return
	}
	Fprint(w, string(jsonData))

	Println("----  get_all() ends  ----")
}

func get(w http.ResponseWriter, r *http.Request) {
	Println("---- get() begins ----")

	r.ParseForm()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := sql.Open("sqlite3", "comento/comento.db")
	if err != nil {
		Println("Failed to open the database file: ", err)
		return
	}

	rows, err := db.Query(`SELECT name, comment FROM Comento WHERE thread = ?;`,
			r.Form["thread"][0])
	if err != nil {
		Println("Failed to get data: ", err)
		return
	}
	defer rows.Close()

	var data = []Data{}
	for rows.Next() {
		var name string
		var comment string

		err := rows.Scan(&name, &comment)
		if err != nil {
			Println("Failed to scan data: ", err)
			return
		}

		data = append(data, Data{name, comment})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		Println("Failed to marshal data: ", err)
		return
	}
	Fprint(w, string(jsonData))

	Println("----  get() ends  ----")
}

func post(w http.ResponseWriter, r *http.Request) {
	Println("---- post() begins ----")

	r.ParseForm()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	Fprintln(w, r.Form)
	Fprintln(w, r.Form["thread"][0])

	db, err := sql.Open("sqlite3", "comento/comento.db")
	if err != nil {
		Println("Failed to open the database file: ", err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Comento
			(id      INTEGER,
			 thread  TEXT,
			 name    TEXT,
			 comment TEXT,
		 	 PRIMARY KEY (id));`,
	)
	if err != nil {
		Println("Failed to create a table: ", err)
		return
	}

	_, err = db.Exec(`INSERT INTO Comento (thread, name, comment) VALUES (?, ?, ?)`,
		r.Form["thread"][0], r.Form["name"][0], r.Form["comment"][0])
	if err != nil {
		Println("Failed to insert a record: ", err)
	}

	Println("----  post() ends  ----")
}

func main() {
	Println("---- Comento begins ----")

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/get_all", get_all)
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)

	err := server.ListenAndServe()
	if err != nil {
		Println("ListenAndServe() error: ", err)
	}

	Println("----  Comento ends  ----")
}
