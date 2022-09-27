package data

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:sameer@tcp(localhost:3306)/todo?parseTime=true"
var err error

type Task struct {
	// gorm.Model
	Id   int64 `gorm:"primaryKey;autoIncrement:true;unique"`
	Todo string
	Mark int64
}

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed")
	}
}
func CreateTable() {
	Database.AutoMigrate(&Task{})

}
func Mark(id int64) {
	var task Task
	Database.First(&task, id)

}
func CreateTask(id int64, todo string, mark int64) {

	us := Task{Id: id, Todo: todo, Mark: mark}
	Database.Create(&us)
	// return us
}
func DeleteTask(id int64) {
	user := Task{}
	// Database.Delete(&user,id)
	Database.Where("id = ?", id).Delete(&user)
}

func Test2(id int64) Task {
	user := Task{}
	//  db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result)
	result := Database.Table("tasks").Select("id,todo,mark").Where("id = ?", id)
	// result := Database.First(user,id)
	result.Scan(&user)
	return user
}
func GetAllTasks() ([]Task, int64) {
	var users Task
	result := Database.Find(&users)

	log.Printf("%d rows found.", result.RowsAffected)
	rows, err := result.Rows()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var groups []Task
	for rows.Next() {
		group := new(Task)
		err := rows.Scan(&group.Id, &group.Todo, &group.Mark)
		if err != nil {
			log.Fatalln(err)
		}
		// log.Printf("%+v\n", group[0].Todo)
		groups = append(groups, *group)
	}
	return groups, result.RowsAffected

}

func Test() ([]Task, int64) {
	var users []Task
	result := Database.Find(&users)
	log.Printf("%d rows found.", result.RowsAffected)
	rows, err := result.Rows()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var group []Task

	for rows.Next() {
		err := result.ScanRows(rows, &group)
		if err != nil {
			log.Fatalln(err)
		}

	}

	return group, result.RowsAffected
}
func GetRecord() (uint, string, int64) {
	user := Task{}
	results := Database.Last(&user)
	log.Printf("%d rows found.", results.RowsAffected)
	rows, err := results.Rows()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var group []Task
	for rows.Next() {

		err := results.ScanRows(rows, &group)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", group[0].Todo)

	}
	return uint(group[0].Id), group[0].Todo, group[0].Mark
}
