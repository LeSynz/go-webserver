package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is starting on http://localhost:8080")
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/page.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./static/index.html")
	})



	http.ListenAndServe(":8080", nil)
}