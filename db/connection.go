package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/moisesPompilio/api-rest-go/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = conn.Ping()

	return conn, err

}
