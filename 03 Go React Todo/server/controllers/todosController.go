package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/radhika.parmar/go-react-todo/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetTodos()

	w.Header().Set("Content-Type", "application/json")
	if err.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo

	// need to unmarshal body
	w.Header().Set("Content-Type", "application/json")
	if body, err := io.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(body, &todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to read request body"})
		return
	}

	// need to validate
	validate := validator.New()
	validationError := validate.Struct(todo)
	if validationError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors := make(map[string]string)

		for _, value := range validationError.(validator.ValidationErrors) {
			errors[value.Field()] = value.ActualTag()
		}
		return
	}

	// need to create data
	result, err := todo.CreateTodo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"errors": err.Error()})
		return
	}

	// return statements by setting headers and response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := strconv.ParseInt(params["todoId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid todo id"})
		return
	}

	todos, _ := models.DeleteTodo(id)
	if todos != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong"})
	}

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	w.Header().Set("Content-Type", "application/json")
	// check json body
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, &todo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unable to read request body"})
			return
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json body"})
		return
	}
	// get todo by id
	var params = mux.Vars(r)
	id, err := strconv.ParseInt(params["todoId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid todo ID"})
		return
	}

	getTodo, _ := models.GetTodoById(id)

	// update data
	if todo.Completed == true || todo.Completed == false {
		getTodo.Completed = todo.Completed
	}
	if todo.Title != "" {
		getTodo.Title = todo.Title
	}
	models.SaveResult(getTodo)

	//return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getTodo)

}
