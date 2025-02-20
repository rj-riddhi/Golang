package models

import (
	"strconv"

	"strings"

	"githlab.com/radhika.parmar/go-jwt-auth-project/database"
	"gorm.io/gorm"
)

type User struct {
	ID            int    `json:"id"`
	First_name    string `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     string `json:"last_name" validate:"required,min=2,max=100"`
	Password      string `json:"password" validate:"required,min=6"`
	Email         string `json:"email" validate:"email,required" gorm:"unique"`
	Phone         string `json:"phone" validate:"required"`
	Token         string `json:"token"`
	User_type     string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token string `json:"refresh_token"`
	User_id       string `json:"user_id"`
}

type PaginatedResponse struct {
	Users         []User `json:"users"`
	TotalRecords  int    `json:"totalRecords"`
	TotalPages    int    `json:"totalPages"`
	CurrentPage   int    `json:"currentPage"`
	RecordPerPage int    `json:"recordPerPage"`
}

var db *gorm.DB

func init() {
	db = database.Connect()
	db.AutoMigrate(&User{})
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Id = ?", Id).Find(&getUser)
	return &getUser, db
}

func (user *User) CreateUser() (*User, error) {
	result := db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	user.User_id = strconv.Itoa(user.ID)
	db.Save(&user)
	return user, nil
}

func GetAllUser() ([]User, *gorm.DB) {
	var users []User
	err := db.Find(&users)
	return users, err
}

func GetAllUserPaginated(limit int, offset int, search string) ([]User, int64, int64, error) {
	var users []User
	var (
		totalRecords int64
		totalPage    int64
	)
	db.Find(&users).Count(&totalRecords)
	totalPage = int64((totalRecords + int64(limit) - 1) / int64(limit))
	var result *gorm.DB
	if search != "" {
		searchTerm := "%" + strings.TrimSpace(search) + "%"
		result = db.Where("First_name ILIKE ? OR email ILIKE ?", searchTerm, searchTerm).Limit(limit).Offset(offset).Find(&users)
	} else {
		result = db.Limit(limit).Offset(offset).Find(&users)
	}
	if result.Error != nil {
		return nil, 0, 0, result.Error
	}
	return users, totalRecords, totalPage, nil
}
