package controllers

import (
	// "bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"log"

	"io"
	"net/http"
	"strconv"

	"githlab.com/radhika.parmar/go-jwt-auth-project/database"
	helper "githlab.com/radhika.parmar/go-jwt-auth-project/helpers"
	"githlab.com/radhika.parmar/go-jwt-auth-project/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("Email of password is incorrect")
		check = false
	}
	return check, msg
}

// using gin package/framework
// func UserSignup(c *gin.Context) {
// 	var user models.User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	validationErr := validate.Struct(user)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 		return
// 	}
// Add logic to save the user in the database, etc.
// 	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully!"})
// }

// using normal http and mux package
func UserSignup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user = ReadBody(w, r, user)

	// Validate the unmarshaled user object
	validationErr := validate.Struct(user)
	if validationErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Prepare a map to hold the validation errors
		errors := make(map[string]string)
		for _, err := range validationErr.(validator.ValidationErrors) {
			// Use the field name as the key and the error message as the value
			errors[err.Field()] = err.ActualTag() + " validation failed"
		}

		// Encode the errors map to JSON and send as a response
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Validation failed",
			"errors":  errors,
		})
		return
	}

	db := database.Connect()
	var count int64
	db.Where("email = ? OR Phone = ?", user.Email, user.Phone).Find(&user).Count(&count)

	if count > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "This Email or Phonenumber exist"})
		return
	}
	password := HashPassword(user.Password)
	user.Password = password

	// to set the tokens
	token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)
	user.Token = token
	user.Refresh_token = refreshToken
	// If validation passed, create the user
	newUser, err := user.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create user: " + err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	var user models.User
	user = ReadBody(w, r, user)

	var foundUser models.User
	err := db.Where("Email = ?", user.Email).Find(&foundUser)
	if err.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error)
		return
	}

	is_valid, message := VerifyPassword(user.Password, foundUser.Password)
	if is_valid != true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": message})
		return
	}

	if foundUser.ID == 0 || foundUser.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Record Not found"})
		return
	} else {
		token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)
		// Update tokens in the database
		if err := helper.UpdateAllTokens(db, token, refreshToken, foundUser.ID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update tokens"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// we need update response of found user
		err = db.Where("Email = ?", foundUser.Email).Find(&foundUser)
		if err.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(foundUser)
			return
		}
	}
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	// only admin can access all the users
	w.Header().Set("Content-Type", "application/json")
	r.URL.Query().Set("user_type", w.Header().Get("user_type"))
	if err := helper.CheckUserType(w, "ADMIN"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"errors": err.Error()})
		return
	}

	// to add pagination
	recordPerPage, err := strconv.Atoi(r.URL.Query().Get("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}
	page, err1 := strconv.Atoi(r.URL.Query().Get("page"))
	if err1 != nil || page < 1 {
		page = 1
	}
	startIndex := (page - 1) * recordPerPage

	// need to add filters
	search := r.URL.Query().Get("search")
	users, totalRecords, totalPages, _ := models.GetAllUserPaginated(recordPerPage, startIndex, search)
	//add metadata of pagination
	response := models.PaginatedResponse{
		Users:         users,
		TotalRecords:  int(totalRecords),
		TotalPages:    int(totalPages),
		CurrentPage:   page,
		RecordPerPage: recordPerPage,
	}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// only admin have access to the open any user, regular user can not access it
	if err := helper.MatchUserTypeToUid(w, userId); err != nil {
		fmt.Println("error while parsing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := models.GetUserById(Id)

	res, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "applicatio/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ReadBody(w http.ResponseWriter, r *http.Request, user models.User) models.User {
	// Read the request body and unmarshal it into the user object
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, &user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return user
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to read request body"})
		return user
	}
	return user
}
