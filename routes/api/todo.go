package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
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
			var updatedTodo Todo
			json.NewDecoder(r.Body).Decode(&updatedTodo)
			for i, t := range todos {
				if t.ID == updatedTodo.ID {
					todos[i] = updatedTodo
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			http.Error(w, "Todo not found", http.StatusNotFound)
		case "DELETE":
			idParam := r.URL.Query().Get("id")
			if idParam == "" {
				http.Error(w, "Missing id parameter", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "Invalid id parameter", http.StatusBadRequest)
				return
			}
			for i, t := range todos {
				if t.ID == id {
					todos = append(todos[:i], todos[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.Error(w, "Todo not found", http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
		}
	})
}