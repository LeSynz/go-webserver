package main

import (
	"fmt"
	"go-webserver/routes"
	"net/http"
)

func main() {
    fmt.Println("Server is starting on http://localhost:8080")

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    routes.RegisterMainRoutes()
    routes.RegisterApiRoutes()

    http.ListenAndServe(":8080", nil)
}