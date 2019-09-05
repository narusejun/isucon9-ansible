package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// Book ...
type Book struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	BorrowerID int    `db:"borrowerId"`
}

// User ...
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

const rs2Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString2(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rs2Letters[rand.Intn(len(rs2Letters))]
	}
	return string(b)
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
	schema := `CREATE TABLE IF NOT EXISTS users (
		id integer not null primary key,
		name text,
		age integer);`

	result, err := db.Exec(schema)

	if err != nil {
		log.Printf("Error: %q", err)
	} else {
		log.Printf("Result: %q", result)
	}
}

// db.Exec
func createBook() {
	schema := `CREATE TABLE IF NOT EXISTS books (
		id integer not null primary key,
		name text,
		borrowerId integer);`

	result, err := db.Exec(schema)

	if err != nil {
		log.Printf("Error: %q", err)
	} else {
		log.Printf("Result: %q", result)
	}
}

// db.Queryx
func selectUserQueryx() {
	sentence := `SELECT * FROM users;`

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
	sentence := `SELECT * FROM users;`

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

func prepareNPlus1() {
	for i := 0; i < 20; i++ {
		id := rand.Intn(1145141919)
		name := randString2(5)
		age := rand.Intn(100)

		for true {
			_, err := db.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?);", id, name, age)
			if err == nil {
				break
			}
		}

		log.Printf("%d, %q, %d", id, name, age)

		bollowerID := id
		id = rand.Intn(1145141919)
		name = randString2(5)

		for true {
			_, err := db.Exec("INSERT INTO books (id, name, borrowerId) VALUES (?, ?, ?);", id, name, bollowerID)
			if err == nil {
				break
			}
			log.Println(err)
		}

		log.Printf("%d, %q, %d", id, name, bollowerID)
	}

}

func nPlus1() {
	sentence := `SELECT * FROM books;`

	bs := []Book{}

	err := db.Select(&bs, sentence)

	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	for _, b := range bs {
		sent := `SELECT * FROM users WHERE id=?;`
		u := User{}

		err := db.Get(&u, sent, b.BorrowerID)

		if err != nil {
			log.Println(err)
			continue
		}

		// fmt.Println("SUCCESS!: ", b, u)
	}
}

func nPlus1EagerLoading() {
	sentenceBooks := `SELECT * FROM books;`

	bs := []Book{}
	err := db.Select(&bs, sentenceBooks)
	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	sentenceUsers := `SELECT * FROM users WHERE id IN (%s);`

	ids := []string{}

	for _, b := range bs {
		ids = append(ids, strconv.Itoa(b.BorrowerID))
	}

	sentenceUsers = fmt.Sprintf(sentenceUsers, strings.Join(ids, ", "))

	// fmt.Println(sentenceUsers)

	us := []User{}
	err = db.Select(&us, sentenceUsers)
	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	// for _, u := range us {
	// 	fmt.Println("SUCCESS!: ", u)
	// }
}

func nPlus1Join() {
	sentenceBooks := `SELECT * FROM books;`

	bs := []Book{}
	err := db.Select(&bs, sentenceBooks)
	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	sentenceUsers :=
		`SELECT users.id, users.name, users.age FROM books JOIN users on books.borrowerId=users.id;`

	// fmt.Println(sentenceUsers)

	us := []User{}
	err = db.Select(&us, sentenceUsers)
	if err != nil {
		log.Printf("Error: %q", err)
		return
	}

	// for _, u := range us {
	// 	fmt.Println("SUCCESS!: ", u)
	// }
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func do(exec func()) {
	start := time.Now()

	exec()

	end := time.Now()

	fmt.Println("Using: ", getFunctionName(exec), ", Time: ", end.Sub(start))
}

func main() {
	createUser()
	createBook()

	prepareNPlus1()

	// selectUserSelect()
	// selectUserQueryx()

	do(nPlus1)
	do(nPlus1EagerLoading)
	do(nPlus1Join)
}
