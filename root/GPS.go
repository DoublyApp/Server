package main

import (
        "net/http"
        "database/sql"
        "log"
        _ "github.com/go-sql-driver/mysql"
        "strings"
        "strconv"
)

// /update_gps/<userID>/<latitude>/<longitude>
func UpdateGPS(w http.ResponseWriter, r *http.Request) string {
        db, err_open := sql.Open("mysql", "doubly_app:doubly_user1@tcp(doublydb.ct2fpvea2u25.us-west-2.rds.amazonaws.com:3306)/Doubly")
        if err_open != nil {
                log.Fatal(err_open)
		return strconv.Itoa(0)
        }
        var rStrings = strings.Split(r.URL.Path, "/")
        var rUserID = rStrings[2]
        var rLatitude = rStrings[3]
	var rLongitude = rStrings[4]
	
	var sqlA = "SELECT UserID FROM GPS WHERE UserID = " + rUserID
	rows1, err_check := db.Query(sqlA)
	defer rows1.Close()
	if err_check != nil {
		panic(err_check.Error())
		return strconv.Itoa(0)
	}
	var count = 0
	for rows1.Next() {
		count++
	}

	if count == 0 {
		// insert
		sqlA = "INSERT INTO GPS (UserID, Latitude, Longitude) VALUES (" + rUserID + ", " + rLatitude + ", " + rLongitude + ")"
	}else{
		// update
		sqlA = "UPDATE GPS SET Latitude = " + rLatitude + ", Longitude = " + rLongitude + " WHERE UserID = " + rUserID
	}
	_, err_update := db.Exec(sqlA)
	if err_update != nil {
		panic(err_update.Error())
		return strconv.Itoa(0)
	}
	return strconv.Itoa(1)
}
