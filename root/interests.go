package main

import (
        "net/http"
        "database/sql"
        "log"
        //"strconv"
        _ "github.com/go-sql-driver/mysql"
)

func GetInterests(w http.ResponseWriter, r *http.Request) string {
	db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
        }
        rows, err_query := db.Query("SELECT * FROM Interests")
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
