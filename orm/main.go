package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1221"
	dbname   = "db_1"
)

func main() {
	fmt.Println("")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("dbtabase Connected...")

	defer db.Close()

	// tb := `CREATE TABLE tb_1 [IF NOT EXISTS] (
	// 	 uid int,
	// 	 name varchar(40),
	// 	 age int
	// 	)`

	// err := db.Exec(tb)
	// 	ALTER TABLE table_name
	// ADD COLUMN new_column_name data_type;

	e := true
	for e {

		fmt.Println("Enter 1 to Vew     Data")
		fmt.Println("Enter 2 to Insert  data")
		fmt.Println("Enter 3 to Update  data")
		fmt.Println("Enter 4 to delete  data")
		fmt.Println("Enter 5 to ADD     NEW-Column")
		fmt.Println("Enter 6 to Exit    Programme!")
		var ch int
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Printf("\n")
			vewdata(db)
			fmt.Printf("\n")
		case 2:
			fmt.Printf("\n")
			insertdata(db)
			fmt.Printf("\n")
		case 3:
			fmt.Printf("\n")
			updatetable(db)
			fmt.Printf("\n")
		case 4:
			fmt.Printf("\n")
			deletedata(db)
			fmt.Printf("\n")
		case 5:
			fmt.Printf("\n")
			addNewColumn(db)
			fmt.Printf("\n")
		case 6:
			e = false
		default:
			fmt.Printf("\n")
			fmt.Println("Enter Write Choice..")
			fmt.Printf("\n")
		}

	}

}
func vewdata(db *sql.DB) {

	query := "SELECT * FROM tb_1;"

	rows, e := db.Query(query)

	if e != nil {
		panic(e)
	}

	defer rows.Close()
	fmt.Println("uid  name   age  addr")
	fmt.Println("")
	for rows.Next() {
		var uid string
		var name string
		var age string
		var addr string

		err := rows.Scan(&uid, &name, &age, &addr)

		if err != nil {
			panic(err)
		}
		fmt.Println(uid, " ", name, " ", age, "", addr)

	}
}

func insertdata(db *sql.DB) {

	fmt.Println("Enter Name To be inserted...")
	var n string
	fmt.Scan(&n)
	fmt.Println("Enter age To be inserted...")
	var a string
	fmt.Scan(&a)
	fmt.Println("Enter address To be inserted...")
	var ad string
	fmt.Scan(&ad)

	// 	ins := `
	// INSERT INTO tb_1 (name,age)
	// VALUES ('Biaml', 29)`

	ins := fmt.Sprintf("INSERT INTO tb_1 (name, age, addr) VALUES ('%s','%s','%s') ;", n, a, ad)
	_, e := db.Exec(ins)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Data: name %v , age %v and addr %v Inserted...\n", n, a, ad)

}

func updatetable(db *sql.DB) {

	fmt.Println("Enter uid to update")
	var u string
	fmt.Scan(&u)
	fmt.Println("Enter name To be Updated..")
	var n string
	fmt.Scan(&n)
	fmt.Println("Enter age To be Updated..")
	var a string
	fmt.Scan(&a)
	fmt.Println("Enter addres To be Updated..")
	var ad string
	fmt.Scan(&ad)

	// 	ins := `
	// INSERT INTO tb_1 (name,age)
	// VALUES ('Biaml', 29)`

	ins := fmt.Sprintf("UPDATE tb_1 SET name= '%s', age= '%s', addr= '%s' WHERE uid = '%s' ", n, a, ad, u)
	_, e := db.Exec(ins)
	if e != nil {
		fmt.Println("Error from update func: ", e)
	}
	fmt.Printf("Data: name %v and age %v Updated to the uid  %v addr %v...\n", n, a, u, ad)
}

func deletedata(db *sql.DB) {

	fmt.Println("Enter uid to to delete")
	var u string
	fmt.Scan(&u)
	fmt.Println("Also Enter name To be delete..")
	var n string
	fmt.Scan(&n)

	ins := fmt.Sprintf("DELETE FROM tb_1 WHERE uid=%s AND name='%s'", u, n)
	_, e := db.Exec(ins)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Deleted the uid %v and name %v \n", u, n)
}

func addNewColumn(db *sql.DB) {

	fmt.Println("Enter New Column Name")
	var cl string
	fmt.Scan(&cl)
	fmt.Println("Enter Cilumn Type..")
	var ct string
	fmt.Scan(&ct)

	ins := fmt.Sprintf("ALTER TABLE tb_1  ADD COLUMN %s %s ;", cl, ct)
	_, e := db.Exec(ins)
	if e != nil {
		fmt.Println("Error is:", e)

	}
	fmt.Printf("New Column %v Added !", cl)

}
