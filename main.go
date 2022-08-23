package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1213"
	dbname   = "db_1"
)

type Data struct {
	//gorm.Model
	uid  int
	name string
	age  int
	addr string
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Connected..")
	fmt.Println("DB: ", db)
	db.Create(Data{uid: 12, name: "milan2", age: 22, addr: "jgm"})

}
