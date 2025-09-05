package api

import "net/http"

func RegisterTest() {
    http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("API Test Route"))
    })
}