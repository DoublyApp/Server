package main

import (
        "net/http"
        "database/sql"
        "log"
        _ "github.com/go-sql-driver/mysql"
        "strings"
        "strconv"
)

// /insert_message/<from_user_id>/<to_user_id>/<message text>
func InsertMessage(w http.ResponseWriter, r *http.Request) string {
        db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
        var rStrings = strings.Split(r.URL.Path, "/")
        var fromUser = rStrings[2]
        var toUser = rStrings[3]
        var messageText = rStrings[4]
	messageText = strings.Replace(messageText, "/&", " ", -1)
	var sqlA = "INSERT INTO Messages (SenderID, ReceiverID, MessageText, TimeCreated) VALUES (" + fromUser + ", " + toUser + ", '" + messageText + "', NOW())"
	results, err_query := db.Exec(sqlA)
	if err_query != nil {
		panic(err_query.Error())
		println("Error: could not insert to DB")
		return "{\"Error\":\"Could not insert to DB\"}"
	}
	lastInsertedID, err_last_id := results.LastInsertId()
	if err_last_id != nil {
		panic(err_last_id.Error())
		println("Error: MessageID not found")
                return "{\"Error\":\"messageID not found\"}"
	}
	
	sqlA = "SELECT TimeCreated FROM Messages WHERE MessageID = " + strconv.FormatInt(lastInsertedID, 10)
	rowsQuery, err_query1 := db.Query(sqlA)
	if err_query1 != nil {
		panic(err_query.Error())
		println("Error: MessageID not found")
		return "{\"Error\":\"messageID not found\"}"
	}
	msg := Message{}
	for rowsQuery.Next() {
		msg.MessageID = int(lastInsertedID)
		rowsQuery.Scan(&msg.TimeCreated)
	}
	defer rowsQuery.Close()
	return "{\"MessageID\":" + strconv.FormatInt(lastInsertedID, 10) + ", \"TimeCreated\":\"" + msg.TimeCreated + "\"}"
}
