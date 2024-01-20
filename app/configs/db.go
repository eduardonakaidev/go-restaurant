package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	 Conn *sql.DB
	 err error
)
func OpenConnection() (*sql.DB, error) {
	conf := GetDB()
	
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	Conn, err = sql.Open("postgres", sc)
	if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }
    fmt.Println("Database connected successfully !")
	return Conn, err
}
