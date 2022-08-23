package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type tb_1 struct {
	// ID   int    `gorm:"primaryKey"`
	ID   int    `gorm:"primaryKey;autoIncrement:true;column:id"`
	Uid  int    `gorm:"column:uid"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
	Addr string `gorm:"column:addr"`
	// created_at time.Time
	// updated_at time.Time
	// deleted_at time.Time
}

// type users struct {
// 	Uid  int
// 	Name string
// 	Age  int
// 	Addr string
// }

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1213"
		dbname   = "db_1"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Db Connected...")
	}

	db.AutoMigrate(&tb_1{})
	// db.Create(&tb_1{Uid: 6, Name: "Nirmal", Age: 25, Addr: "Kolkata"})
	// db.Create(&tb_1{Uid: 9, Name: "Bolai", Age: 21, Addr: "Gopalgonj"})
	// db.Create(&tb_1{Uid: 66, Name: "Madhobda", Age: 22, Addr: "Kolkata"})
	// db.Create(&tb_1{Uid: 55, Name: "Bolai", Age: 27, Addr: "Jhargram"})

	fmt.Printf("DB %T \n", db)
	// u := tb_1{Uid: 33, Name: "nimai", Age: 20, Addr: "Gokuldham"}
	// db.Create(&u)

	// db.First(&user)

	// er := db.Migrator().AddColumn(&tb_1{}, "uid")
	// if er != nil {
	// 	fmt.Println(er)
	// }

	// db.Migrator().AddColumn(&tb_1{}, "name")
	// db.Migrator().AddColumn(&tb_1{}, "age")
	// db.Migrator().AddColumn(&tb_1{}, "addr")

	// Create

	// person := tb_1{name: "Jinzhu", age: 18, addr: "jgm"}

	// fmt.Println("res:", result)

	// result := db.Create(&person)
	// fmt.Println("res:", result)

	// Read
	// var p Data
	// db.First(&p, 1)             // find product with integer primary key
	// db.First(&p, "uid =?", "1") // find product with code D42

	// Update - update
	// db.Model(&p).Update("uid", 1)
	// // Update - update multiple fields
	// db.Model(&p).Updates(Data{uid: 2, name: "arijit"}) // non-zero fields
	// db.Model(&p).Updates(map[string]interface{}{"name": "laltu", "uid": 1})

	// // Delete - delete product
	// db.Delete(&p, 1)
	// var us tb_1
	// db.Where("ID = ?", 4).Find(&us)
	// us.Name = "Bisakha"
	// us.Age = 100
	// us.Uid = 24
	// us.Addr = "jhargram"
	// db.Save(&us)
	db.Model(&tb_1{}).Where("id = ?", 10).Update("name", "Rikisi")
	db.Model(&tb_1{}).Where("id = ?", 10).Update("age", 200)

	var dt tb_1
	er := db.Model(&tb_1{}).Where("id = ?", 9).Delete(dt)
	if er != nil {
		fmt.Println("Er on delete->", er)
	}

	var ud []tb_1
	db.Find(&ud)
	fmt.Println("DATAS FROM DATBASE ARE ......")
	for _, v := range ud {
		// if v.ID == 2 {
		// 	fmt.Printf("ID : %v, Uid : %v, Name : %v, Age : %v, Addr: %v \n", v.ID, v.Uid, v.Name, v.Age, v.Addr)
		// }
		fmt.Printf("ID : %v, Uid : %v, Name : %v, Age : %v, Addr: %v \n", v.ID, v.Uid, v.Name, v.Age, v.Addr)

	}
	// var dt tb_1
	// db.Model(&dt).Updates(tb_1{Uid: 5, Name: "Malay", Age: 18, Addr: "jhargram"})
	// db.Model(&dt).Select("Name", "Age").Where("id = ?", 4).Updates(tb_1{Name: "mitali", Age: 20})
	// fmt.Println("Updated..")

	// db.Where("name = ?", "Bisakha").Delete(&dt)

	// drop colum not tested yet
	//db.Model(&tb_1{}).Migrator().DropColumn(&tb_1{},"uid")
}
