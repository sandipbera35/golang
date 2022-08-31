package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Marksheet struct {
	Id        int    `gorm:"primaryKey:true;autoIncrement:true;column:id"`
	Subject   string `gorm:"column:subject"`
	Marks     int    `gorm:"column:marks"`
	StudentID int
}

type Student struct {
	ID        int       `gorm:"primaryKey:true;autoIncrement:true;column:id"`
	Name      string    `gorm:"column:name"`
	Age       int       `gorm:"column:age"`
	Addr      string    `gorm:"column:addr"`
	Examtype  string    `gorm:"column:examtype"`
	Marksheet Marksheet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Result struct {
	Name  string
	Total int
	Count int
}

// Student {
// 	id 		  int autoincrement
// 	name 	  string
// 	Marksheet Marksheet
// }

// Marksheet {
// 	id 			int autoincrement
// 	marks 		int
// 	StudentID   int
// }

type Search struct {
	ID         int
	Name       string
	Age        int
	Addr       string
	Examtype   string
	Id         int
	Student_id int
	Subject    string
	Marks      int
}

// Retrieve user list with edger loading credit cards
// func GetAll(db *gorm.DB) ([]Marksheet, error) {
//     var users []Marksheet
//     err := db.Model(&Marksheet{}).Preload("Students").Find(&users).Error
//     return users, err
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Db Connected...")
	}

	db.AutoMigrate(&Student{})

	db.AutoMigrate(&Marksheet{})

	// db.Model(&student{}).

	//db.Model(&marksheet{}).("student_id", "student(id)", "RESTRICT", "RESTRICT")
	// db.Model(&marksheet{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")
	chk := true

	for chk {

		fmt.Println("Enter 1 to add Student data ...")

		fmt.Println("Enter 2 to delete... with name")
		fmt.Println("Enter 3 for Search With Name..")
		fmt.Println("Enter 4 to Vew All the Data From Two Table..")
		fmt.Println("Enter 5 to Exit..")
		fmt.Println("Enter 6 to search with GORM")
		fmt.Println("ENter 7 to search name..")
		fmt.Println("Enter 9 to add marksheet data")
		fmt.Println("Enter 12 to get Total marks")
		//fmt.Println("Enter 11 to get First Second Third..")

		var ch int

		fmt.Scan(&ch)

		switch ch {
		case 1:

			fmt.Println("Enter name :")
			var nm string
			fmt.Scan(&nm)

			fmt.Println("Enter age :")
			var ag int
			fmt.Scan(&ag)

			fmt.Println("Enter Address :")
			var addr string
			fmt.Scan(&addr)

			fmt.Println("Enter ExamType :")
			var et string
			fmt.Scan(&et)

			fmt.Println("Enter subject :")
			var sub string
			fmt.Scan(&sub)

			fmt.Println("Enter Student Marks:")
			var sm int
			fmt.Scan(&sm)

			var std Student
			std.Name = nm
			std.Age = ag
			std.Addr = addr
			std.Examtype = et

			var mk Marksheet
			mk.Subject = sub
			mk.Marks = sm
			std.Marksheet = mk

			db.Create(&std)

			//mk.StudentID = std.ID
			//db.Create(&mk)

		case 2:
			fmt.Println("Also Enter name To be delete..")
			var n string
			fmt.Scan(&n)
			var st Student
			// ins := fmt.Sprintf("DELETE FROM students,marksheets WHERE name='%s' AND studets.id = marksheets.student_id", n)
			// db.Exec(ins)
			db.Where("name = ?", n).Delete(&st)

			fmt.Printf("Deleted the name %s \n", n)
		case 3:
			fmt.Println("")

			fmt.Println("enter name to search: ")
			var name1 string
			fmt.Scan(&name1)

			// useaing orm
			// rows, _ := db.Model(&Student{}).Where("name = ?", name1).Select("name, age").Rows()
			// // .Where("name = ?", "nilima")

			// defer rows.Close()

			// for rows.Next() {
			// 	var name string
			// 	var age int

			// 	rows.Scan(&name, &age)
			// 	fmt.Println("Gorm:", "name:", name, "age", age)

			// 	// do something
			// }

			var search Search
			sql1 := fmt.Sprintf(`SELECT students.id,students.name,students.age,students.examtype,
			subject,marks FROM students,marksheets
			where students.id = marksheets.student_id
			AND students.name = '%s';`, name1)

			db.Raw(sql1).Scan(&search)

			if search.ID != 0 {

				fmt.Println("ID:", search.ID, "Name:", search.Name, "Age:", search.Age,
					"ExamType: ", search.Examtype, "Subject: ", search.Subject, "marks:", search.Marks)

				fmt.Println("")
				fmt.Println("")

			} else {
				fmt.Println("")
				fmt.Println("Name not found...")
				fmt.Println("")
			}

		case 4:

			fmt.Println("All the data from Both Table are ....")
			fmt.Println("")
			var search []Search
			sql1 := `SELECT * FROM students,marksheets where students.id = marksheets.student_id;`

			// students.id,students.name,students.age,students.examtype,
			// subject,marks

			db.Raw(sql1).Scan(&search)
			for i := 0; i < len(search); i++ {

				fmt.Println(search[i])

			}
			// fmt.Println("Search", search)

			fmt.Println("")
			fmt.Println("")

		case 5:
			chk = false
		case 6:

			fmt.Println("Enter name to search with GORM: ")
			var name2 string
			fmt.Scan(&name2)

			var st []Student
			//
			//db.Joins("Marksheet").Where("name = ?", name2).Find(&st)

			//var meeting model.Meeting
			db.Preload("Marksheet").Where("name = ?", name2).Find(&st)

			fmt.Println(st)
			fmt.Println("")

			fmt.Println(".........................................................................................")
			fmt.Println("")
			for _, v := range st {
				fmt.Println("ID:", v.ID, "Name: ", v.Name, "Age: ", v.Age, "ExamType: ",
					v.Examtype, "Subject: ", v.Marksheet.Subject, "Marks: ", v.Marksheet.Marks, "Student_id: ", v.Marksheet.StudentID)
			}

			fmt.Println("")
			fmt.Println("")

		case 7:
			fmt.Println("Enter name to search: ")
			var name2 string
			fmt.Scan(&name2)

			var st []Student

			nm := strings.ToLower(name2)
			//
			//db.Joins("Marksheet").Where("name = ?", name2).Find(&st)

			//var meeting model.Meeting
			db.Preload("Marksheet").Where("lower(name) LIKE ?", "%"+nm+"%").Find(&st)

			// fmt.Println(st)
			fmt.Println("")
			var arr []string
			fmt.Println("")
			fmt.Println("")
			for _, v := range st {
				// fmt.Println("ID:", v.ID, "Name: ", v.Name, "Age: ", v.Age, "ExamType: ",
				// 	v.Examtype, "Subject: ", v.Marksheet.Subject, "Marks: ", v.Marksheet.Marks, "Student_id: ", v.Marksheet.StudentID)
				arr = append(arr, v.Name)

			}

			if len(arr) == 0 {
				fmt.Println("No sweetable match found.. with :", name2)
				fmt.Println("")
				fmt.Println("")
				break
			}

			// fmt.Println("arr:", arr)

			//var store []string

			// var arrs []string
			for i := 0; i < len(arr); i++ {
				// arrs = append(arrs, arr[i])
				fmt.Printf("Do you mean '%v' ? (if yes Enter Y / if no enter N) : \n", arr[i])

				//............................................................
				// if arrs == nil {
				// 	fmt.Printf("No match Found as '%v' ..", nm)
				// }

				var cho string
				fmt.Scan(&cho)
				cho1 := strings.ToLower(cho)
				var ul []Student
				if cho1 == "y" {
					// for j := 0; j < len(arr); j++ {

					db.Preload("Marksheet").Where("name = ?", arr[i]).Find(&ul)

					for _, v := range ul {

						fmt.Println("Name: ", v.ID, "Name :", v.Name, "Age : ", v.Age, "examtype :", v.Examtype,
							"Addr: ", v.Addr, "Subject: ", v.Marksheet.Subject, "Student_id: ", v.Marksheet.StudentID)
					}
					// }
					if cho1 == "y" {
						break
					}
				}
				//........................................................
			}
			// fmt.Println("arrs", arrs)

			fmt.Println("")
			fmt.Println("")

		case 8:
		// //takng value in map.....
		// // var results []map[string]interface{}
		// var map1 []map[string]Student
		// //db.Table("students").Find(&results)
		// db.Preload("Marksheet").Find(&map1)
		// // fmt.Println("Results:", results)
		// if len(map1) == 0 {
		// 	fmt.Println("map not initialised..")
		// 	break
		// }
		// for _, v := range map1 {
		// 	fmt.Println(v)
		// }
		// fmt.Println("")
		// fmt.Println("re0: ", map1[0])
		// fmt.Println("")
		case 9:

			var p []Student
			var ar []string
			db.Preload("Marksheet").Find(&p)

			for _, v := range p {
				ar = append(ar, v.Name)
			}
			for i := 0; i < len(ar); i++ {
				fmt.Println("Name : ", ar[i])
			}
			fmt.Println("...........................")
			fmt.Println("")
			fmt.Println("Enter name  marksheet data: ")
			var name3 string
			fmt.Scan(&name3)

			var flag bool

			for i := 0; i < len(ar); i++ {
				if ar[i] == name3 {

					flag = true
					break

				} else {
					flag = false

				}

			}
			if !flag {
				fmt.Println("Name is not valid")

			} else {

				fmt.Println("Enter Subject:  ")
				var sub string
				fmt.Scan(&sub)
				fmt.Println("Enter Marks:  ")
				var mrk int
				fmt.Scan(&mrk)

				var ul Student

				db.Preload("Marksheet").Where("name = ?", name3).Find(&ul)

				vl := ul.ID

				// fmt.Println("Var:", vl)

				var mk Marksheet
				mk.Subject = sub
				mk.Marks = mrk
				mk.StudentID = vl
				ul.Marksheet = mk

				db.Create(&mk)
				fmt.Println("Data Entered...")
			}
		case 10:
			// total marks with student name :

			fmt.Println("Enter student name to get total marks :")
			var nma string
			fmt.Scan(&nma)

			// var stn []Student
			// // var arrs []int

			// db.Preload("Marksheet").Select("sum(marks)").Where("name = ?", nma).Find(&stn)

			var ul Student

			db.Preload("Marksheet").Where("name = ?", nma).Find(&ul)

			vl := ul.ID

			if vl == 0 {
				fmt.Println("")
				fmt.Println(nma, "name is not exist")
				fmt.Println("")
				fmt.Println("Please choose name from below mantion :")
				fmt.Println("")
				var st []Student
				db.Table("students").Find(&st)
				for _, v := range st {
					fmt.Println("Name: ", v.Name)
				}
				fmt.Println("")

				break
			}
			var mkr []Marksheet

			db.Where("student_id = ?", vl).Find(&mkr)
			sumval := 0
			for _, v := range mkr {

				sumval += v.Marks

			}
			fmt.Println("")
			fmt.Println("Total Marks Of ", nma, " is: ", sumval)
			fmt.Println("")

		case 11:

			var st []Student
			db.Preload("Marksheet").Find(&st)
			fmt.Println(len(st))

			for _, v := range st {
				var ul Student
				db.Preload("Marksheet").Where("name = ?", v.Name).Find(&ul)

				vl := ul.ID
				var mkr []Marksheet

				db.Where("student_id = ?", vl).Find(&mkr)
				sumval := 0
				for _, v := range mkr {

					sumval += v.Marks

				}
			}
		case 12:

			fmt.Println("Enter student name to get total marks :")
			var nma string
			fmt.Scan(&nma)

			// with sum with sql

			var ul Student

			db.Preload("Marksheet").Where("name = ?", nma).Find(&ul)

			ids := float64(ul.ID)

			if ids == 0 {
				fmt.Println("name :", nma, " ....not exist !")
				break
			}

			var tot int64
			var vlav float64
			var count int64
			db.Model(&Marksheet{}).Where("student_id = ?", ids).Count(&count)
			// co := int(count)
			fmt.Println("count :", count)
			// fmt.Println("cou : ", co)

			db.Table("marksheets").Select("sum(marks)").Where("student_id = ?", ids).Scan(&tot)
			db.Table("marksheets").Select("avg(marks)").Where("student_id = ?", ids).Scan(&vlav)
			// db.Table("marksheets")
			// summax := 100 * co // vl
			// per := vl / (100 * count)

			fmt.Println("")
			fmt.Println("Total Marks of ", nma, "is: ", tot)
			fmt.Println("")
			fmt.Println("Total Avg marks of ", nma, "is: ", vlav)
			fmt.Println("")
			fmt.Println("percentage is : ", (tot/count*100)/100, "%")
			fmt.Println("")

			//   var results Result
			//   db.Table("marksheets").Joins("students").Where("students_id = ?", vl).
			//   Select("date(created_at) as date, sum(marks) as total").
			//   Group("name(name)").
			//   Having("sum(marks)").Scan(&results)
			// var st2 Student
			// var re Result
			// fmt.Println("")
			// db.Table("marksheets, students").Group("name").Select("name , sum(marks), count(student_id)").Where("name = ?", nma).Scan(&re)
			// fmt.Println("re:", re)
			// fmt.Println("")

		default:

			fmt.Println("")
			fmt.Println("Enter Correct Choice Please...!")
			fmt.Println("")

		}
	}
}
