package api

import "net/http"

func RegisterTest() {
    http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
        // example query idk
        limit := r.URL.Query().Get("limit")
        if limit != "" {
            w.Write([]byte("API Testing Route with limit: " + limit))
            return
        }
        w.Write([]byte("API Testing Route"))
    })
}