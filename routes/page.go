package routes

import "net/http"

func RegisterPage() {
    http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/page.html")
    })
}