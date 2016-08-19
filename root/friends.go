package main

/*
import (
	"io"
	"net/http"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func getFriends(w http.ResponseWriter, r *http.Request, userID int) string {
	io.WriteString(w, "Hello world!")

	db, err := sql.Open("mysql", "root:polkatis4foreverything@/Doubly")
  if err := db.Ping(); err != nil {
    log.Fatal(err)
  }
	rows, err := db.Query("SELECT * FROM UsersFriends WHERE UserID = " + string(userID))
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s is %d\n", user.UserName, age)
		io.WriteString(w, name);
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return formatFriends(friends)
}
*/
