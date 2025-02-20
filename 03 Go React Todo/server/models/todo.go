package models

import (
	// "fmt"

	database "github.com/radhika.parmar/go-react-todo/databse"
	"gorm.io/gorm"
)

type ToDo struct {
	ID        uint   `json:"Id"`
	Title     string `json:"title" validation:"required"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB

func init() {
	db = database.Connect()
	db.AutoMigrate(&ToDo{})
}

func GetTodos() ([]ToDo, *gorm.DB) {
	var todos []ToDo
	err := db.Find(&todos)
	return todos, err
}

func (todo ToDo) CreateTodo() (*ToDo, error) {
	result := db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func DeleteTodo(id int64) ([]ToDo, error) {
	var todos []ToDo
	if err := db.Delete(&ToDo{}, id).Error; err != nil {
		return nil, nil
	}

	if err := db.Find(&todos).Error; err != nil {
		return nil, nil
	}
	return todos, nil
}

func GetTodoById(id int64) (*ToDo, *gorm.DB) {
	var todo ToDo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, nil
	}
	return &todo, nil
}

func SaveResult(todo *ToDo) {
	db.Save(todo)
	return
}
