package apis

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DBHandle *sql.DB = nil

func OpenDB(ip string, db string) (*sql.DB, error) {
	//dbStr := "dunkin:dunkin123@tcp(10.204.166.227:3306)/mufg_dec1?multiStatements=true"
	//db := "mufg_dec1"

	dbStr := fmt.Sprintf("root:my-secret-pw@tcp(%s:3306)/%s?multiStatements=true", ip, db)

	fmt.Println(dbStr)

	DBH, err := sql.Open("mysql", dbStr)
	if err != nil {
		panic(err)
	}

	DBHandle = DBH

	//defer DBHandle.Close()

	// See "Important settings" section.
	DBHandle.SetConnMaxLifetime(time.Minute * 3)
	DBHandle.SetMaxOpenConns(10)
	DBHandle.SetMaxIdleConns(10)

	err = DBHandle.Ping()
	if err != nil {
		DBHandle.Close()
		return nil, err
	}

	return DBHandle, nil
}

type User struct {
    ID     int64
	create_time string
    name  string
    password string
}


// albumsByArtist queries for albums that have the specified artist name.
func GetUserInfo(db *sql.DB) ([]User, error) {
    // An albums slice to hold data from returned rows.
	fmt.Print("Inside GetUserInfo\n\n")
    var userList []User

    rows, err := db.Query("SELECT * FROM user_info");
    if err != nil {
        return nil, fmt.Errorf("getUserInfo %q", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
	
    for rows.Next() {
		
        var alb User
        if err := rows.Scan(&alb.ID, &alb.create_time, &alb.name, &alb.password); err != nil {
            return nil, fmt.Errorf("getUserInfo %q", err)
        }
        userList = append(userList, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getUserInfo %q", err)
    }
    return userList, nil
}

func DisplayUsers(u []User) {
	for _, user := range u {
		fmt.Printf("ID: %d, Name: %s, Password: %s\n", user.ID, user.name, user.password)
	}
}

