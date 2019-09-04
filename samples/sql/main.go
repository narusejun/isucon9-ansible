package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// User ...
type User struct {
	Name string `json:"name" db:"name"`
	Age  int64  `json:"age" db:"age"`
}

func init() {
	seedBuf := make([]byte, 8)
	crand.Read(seedBuf)
	rand.Seed(int64(binary.LittleEndian.Uint64(seedBuf)))

	dbHost := "127.0.0.1"

	dbPort := "3306"

	dbUser := "root"

	dbPassword := ":" + "114514YJsnpi1919$"

	dsn := fmt.Sprintf("%s%s@tcp(%s:%s)/isubata?parseTime=true&loc=Local&charset=utf8mb4",
		dbUser, dbPassword, dbHost, dbPort)

	log.Printf("Connecting to db: %q", dsn)

	tmpdb, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	db = tmpdb

	for {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(time.Second * 3)
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Printf("Succeeded to connect db.")
}

// db.Exec
func createUser() {
	schema := `CREATE TABLE IF NOT EXISTS user (
		name text,
		age integer);`

	result, err := db.Exec(schema)

	if err != nil {
		log.Printf("Error: %q", err)
	} else {
		log.Printf("Result: %q", result)
	}
}

// db.MustExec
func insertUser() {
	sentence := `INSERT INTO user (name, age) VALUES ("david", 114514);`
	db.MustExec(sentence)
}

// db.Queryx
func selectUserQueryx() {
	sentence := `SELECT * FROM user;`

	rows, err := db.Queryx(sentence)

	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	for rows.Next() {
		var u User
		err = rows.StructScan(&u)

		if err != nil {
			log.Printf("Error: %q", err)
			return
		}

		fmt.Println(u)
	}
}

// db.Select
func selectUserSelect() {
	sentence := `SELECT * FROM user;`

	us := []User{}

	err := db.Select(&us, sentence)

	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	for _, u := range us {
		fmt.Println(u)
	}
}

func nPlus1() {

}

func nPlus1EagerLoading() {

}

func nPlus1Join() {

}

func main() {
	createUser()

	insertUser()

	selectUserQueryx()

	selectUserSelect()
}
