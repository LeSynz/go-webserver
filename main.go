package main

import (
	"fmt"
	"go-webserver/routes"
	"net/http"
	"time"
)

func main() {
    fmt.Println("Server is starting on http://localhost:8080")

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    routes.RegisterMainRoutes()
    routes.RegisterApiRoutes()
    // wait 2s
    time.Sleep(2 * time.Second)
        fmt.Println("Routes registered. Server is ready to handle requests.")

    http.ListenAndServe(":8080", nil)
}