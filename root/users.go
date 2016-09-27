package main

import (
	//"io"
	"net/http"
	"database/sql"
	"log"
	//"fmt"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

// returns userID to client
// /insert_user/<user name>/<email>/<password>/<DOB>/<gender>
func InsertUser(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
	var rStrings = strings.Split(r.URL.Path, "/")
	var rUserName = rStrings[2]
	var rUserEmail = rStrings[3]
	var rPassword = rStrings[4]
	var rDOB = rStrings[5]
	var rGender = rStrings[6]
        rows, err_query := db.Query("SELECT * FROM Users WHERE Users.Email = '" + rUserEmail + "'")
	defer rows.Close()
        if err_query != nil {
                panic(err_query.Error())
        }
	var count = 0
	for rows.Next() {
		count++
	}
	if count > 0 {
		return "{\"Error\":\"UserExists\"}"
	}
	results_insert, err_insert := db.Exec("INSERT INTO Users(UserName, Email, Password, DOB, Gender) VALUES ('" + rUserName + "', '" + rUserEmail + "', '" + rPassword + "', " + rDOB + ", '" + rGender + "')")
	if err_insert != nil {
		panic(err_insert.Error())
	}
	lastInsertedID, err_last_id := results_insert.LastInsertId()
	if err_last_id != nil {
		println("Error: UserID not found")
		panic(err_last_id.Error())
		return "{\"Error\":\"UserID Not Found\"}"
	}
	return "{\"UserID\":\"" + strconv.FormatInt(lastInsertedID, 10) + "\"}"
}

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
