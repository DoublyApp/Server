package main

import (
	"io"
	"net/http"
	"database/sql"
	"log"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) string {
	io.WriteString(w, "Hello world!")

	db, err_open := sql.Open("mysql", "doubly_user:db_user1@/Doubly")
  if err_open != nil {
		log.Fatal(err_open)
  }
	rows, _ := db.Query("SELECT * FROM Users")
	defer rows.Close()
  var users []User
	for rows.Next() {
		user := User{}
		var userID, userName, email, password, dob, gender []byte
		rows.Scan(&userID, &userName, &email, &password, &dob, &gender)
		user.UserID, _ = strconv.Atoi(string(userID))
		user.UserName = string(userName)
		user.Email = string(email)
		user.Password = string(password)
		user.DOB = string(dob)
		user.Gender = string(gender)
		fmt.Printf("%s is %s\n", string(user.UserName), string(user.DOB))
		//io.WriteString(w, user.UserName);
		users = append(users, user)
	}
	//if err := rows.Err(); err != nil {
	//	log.Fatal(err)
	//}

	return FormatUsers(users)
}
