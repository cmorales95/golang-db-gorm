package storage

import (
	"fmt"
	"log"
	"sync"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

//Driver of storage
type Driver string

const (
	Mysql    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

//New create a connection with DB
func New(d Driver) {
	switch d {
	case Mysql:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newPostgresDB singleton
func newPostgresDB() {
	once.Do(func(){
		var err error
		db, err = gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/godb?sslmode=disable"),&gorm.Config{})
		if err != nil {
			log.Fatalf("Connection with DB is not available: %v", err)
		}
		fmt.Println("Connected to Postgres!")
	})
}

// newMySQLDB singleton
func newMySQLDB() {
	once.Do(func(){
		var err error

		db, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/godb?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Connection with DB is not available: %v", err)
		}
		fmt.Println("Connected to MySQL!")
	})
}

// Pool return a unique instance of DB
func DB() *gorm.DB {
	return db
}
