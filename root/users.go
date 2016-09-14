package main

import (
	//"io"
	"net/http"
	"database/sql"
	"log"
	//"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
	if err_open != nil {
		log.Fatal(err_open)
	}
	rows, err_query := db.Query("SELECT * FROM Users")
	if err_query != nil {
		panic(err_query.Error())
	}
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
		users = append(users, user)
	}
	defer rows.Close()
	return FormatUsers(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
        var rID = r.URL.Path[len("/get_user_by_id/"):]
        rows, err_query := db.Query("SELECT * FROM Users WHERE Users.UserID = " + rID)
        if err_query != nil {
                panic(err_query.Error())
        }
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
                users = append(users, user)
        }
        defer rows.Close()
        return FormatUsers(users)
}
