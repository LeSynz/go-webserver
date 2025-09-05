package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

// store the todos in memory
var todos []Todo
var nextID int = 1 // first id

func RegisterTodo() {
	http.HandleFunc("/api/todo", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			// return the todos as json
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos)
		case "POST":
			// create a new todo from the received json
			var newTodo Todo
			if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newTodo.ID = nextID
			nextID++
			todos = append(todos, newTodo)
			log.Printf("Received todo JSON: %+v\n", newTodo)
			w.WriteHeader(http.StatusOK)
		case "PUT":
			w.Write([]byte("PUT Todo Route"))
		case "DELETE":
			w.Write([]byte("DELETE Todo Route"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
		}
	})
}