package main

import (
        "net/http"
        "database/sql"
        "log"
        _ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
)

// /get_interests/<int start>/<int end>/<interest text characters>
func GetInterests(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
	var rStrings = strings.Split(r.URL.Path, "/")
        var rStart = rStrings[2]
        var rEnd = rStrings[3]
	var interestText = rStrings[4]
	var sqlA = "SELECT * FROM Interests "
	if interestText != "" {
		sqlA = sqlA + "WHERE InterestName LIKE('" + interestText + "%')"
	}
        rows, err_query := db.Query(sqlA)
        if err_query != nil {
                panic(err_query.Error())
        }
        var interests []Interest
	startIndex, err_start_conv := strconv.Atoi(rStart)
	if err_start_conv != nil {
		panic(err_start_conv.Error())
		return "{\"Error\":\"invalid START parameters\"}"
	}
	endIndex, err_end_conv := strconv.Atoi(rEnd)
	if err_end_conv != nil {
		panic(err_start_conv.Error())
		return "{\"Error\":\"invalid END parameters\"}"
	}
	var index = 1
        for rows.Next() {
		if index >= startIndex && index <= endIndex {
                	interest := Interest{}
                	rows.Scan(&interest.InterestID, &interest.InterestName)
                	interests = append(interests, interest)
		}else if index > endIndex {
			break
		}
		index++
        }
        defer rows.Close()
        return FormatInterests(interests)
}

func GetUsersInterests(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
	var rUserID = r.URL.Path[len("/get_users_interests/"):]
        rows, err_query := db.Query("SELECT Interests.InterestID, Interests.InterestName FROM UsersInterests INNER JOIN Users ON UsersInterests.UserID = Users.UserID INNER JOIN Interests ON UsersInterests.InterestID = Interests.InterestID WHERE Users.UserID = " + rUserID)
        if err_query != nil {
                panic(err_query.Error())
        }
        var interests []Interest
        for rows.Next() {
                interest := Interest{}
                rows.Scan(&interest.InterestID, &interest.InterestName)
                interests = append(interests, interest)
        }
        defer rows.Close()
        return FormatInterests(interests)
}

// returns interestID to client
// /insert_interest/<userID>/<interestName>
func InsertInterest(w http.ResponseWriter, r *http.Request) string {
	var interestID = -1000
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
        var rStrings = strings.Split(r.URL.Path, "/")
	var rUserID = rStrings[2]
	var rInterestName = rStrings[3]
	println(rInterestName)
	rows, err_query := db.Query("SELECT InterestID FROM Interests WHERE InterestName = '" + rInterestName + "'")
	if err_query != nil {
		panic(err_query.Error())
	}
	var count = 0
	for rows.Next() {
		count++;
		rows.Scan(&interestID)
	}
	if count > 0 {
		// interest exists
		println("Interest Exists " + strconv.Itoa(interestID))
	}else{
		db.Exec("INSERT INTO Interests (InterestName) VALUES (?)", rInterestName)
		rows4, _ := db.Query("SELECT InterestID FROM Interests WHERE InterestName = '" + rInterestName + "'")
		for rows4.Next() {
			rows4.Scan(&interestID)
			println("received interestID: " + strconv.Itoa(interestID))
		}
		defer rows4.Close()
	}
	var query_ui = "SELECT UserID FROM UsersInterests WHERE UserID = " + rUserID + " AND InterestID = " + strconv.Itoa(interestID);
	println(query_ui)
	rows5, err_ui := db.Query("SELECT UserID FROM UsersInterests WHERE UserID = " + rUserID + " AND InterestID = " + strconv.Itoa(interestID))
	if err_ui != nil {
		panic(err_ui.Error())
	}
	count = 0
	for rows5.Next() {
		count++
	}
	if count == 0 {
		db.Exec("INSERT INTO UsersInterests (UserID, InterestID) VALUES (" + string(rUserID) + ", " + strconv.Itoa(interestID) + ")")
		println("UsersInterest already exists")
	}
	defer rows.Close()
	defer rows5.Close()
	return strconv.Itoa(interestID)
}

// /remove_interest/<userID>/<interestID>
// returns 1 if successful remove, 0 if failed to remove
func RemoveInterest(w http.ResponseWriter, r *http.Request) string {
        db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
		return strconv.Itoa(0)
        }
        var rStrings = strings.Split(r.URL.Path, "/")
        var rUserID = rStrings[2]
        var rInterestID = rStrings[3]
	_, err_delete := db.Exec("DELETE FROM UsersInterests WHERE UserID = " + rUserID + " AND InterestID = " + rInterestID)
	if err_delete != nil {
		log.Fatal(err_delete)
		return strconv.Itoa(0)
	}
	return strconv.Itoa(1)
}

// gets the interest's users
// /get_interests_users/<searching user's ID>/<interestID>/<start index>/<end index>
func GetInterestsUsers(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
                return strconv.Itoa(0)
        }
        var rStrings = strings.Split(r.URL.Path, "/")
        // rSearchingUserID will be implemented when GPS searching is active
	//var rSearchingUserID = rStrings[2]
        var rInterestID = rStrings[3]
	var rStartIndex = rStrings[4]
	var rEndIndex = rStrings[5]
	startIndex, err_start_conv := strconv.Atoi(rStartIndex)
        if err_start_conv != nil { 
                panic(err_start_conv.Error())
                return "{\"Error\":\"invalid START parameters\"}"
        }
	endIndex, err_end_conv := strconv.Atoi(rEndIndex)
        if err_end_conv != nil { 
                panic(err_end_conv.Error())
                return "{\"Error\":\"invalid END parameters\"}"
        }
	var sqlA = "SELECT Users.UserID, UserName, DOB, Gender FROM Users INNER JOIN UsersInterests ON Users.UserID = UsersInterests.UserID AND UsersInterests.InterestID = " + rInterestID
	rows1, err_query := db.Query(sqlA)
	if err_query != nil {
		panic(err_query.Error())
	}
	var index = 1
	var users []User
	for rows1.Next() {
		if index >= startIndex && index <= endIndex {
			user := User{}
			rows1.Scan(&user.UserID, &user.UserName, &user.DOB, &user.Gender)
			users = append(users, user)
		}else if index > endIndex {
			break
		}
		index++
	}
	defer rows1.Close()
	return FormatUsers(users)
}
