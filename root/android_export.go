package main

import (
	"io"
	"net/http"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")

	db, err := sql.Open("mysql", "doubly_user:db_user1@/Doubly")
  if err := db.Ping(); err != nil {
    log.Fatal(err)
  }
	println(err)

	io.WriteString(w, getUsers(w, r))
/*
	age := 27
	rows, err := db.Query("SELECT UserName FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
		io.WriteString(w, name);
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
*/
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
