package main

import (
    "log"
    "net/http"
    "strings"
    "time"
)

func logHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
        log.Printf("Completed in %v", time.Since(start))
    })
}

func main() {
    kvStore := &KeyValueStore{}

    mux := http.NewServeMux()

    mux.HandleFunc("/put/", func(w http.ResponseWriter, r *http.Request) {
        key := strings.TrimPrefix(r.URL.Path, "/put/")
        if r.Method == "PUT" {
            putHandler(kvStore, key, w, r)
        } else {
            writeError(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    mux.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
        key := strings.TrimPrefix(r.URL.Path, "/get/")
        if r.Method == "GET" {
            getHandler(kvStore, key, w, r)
        } else {
            writeError(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    mux.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
        key := strings.TrimPrefix(r.URL.Path, "/delete/")
        if r.Method == "DELETE" {
            deleteHandler(kvStore, key, w, r)
        } else {
            writeError(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            listKeysHandler(kvStore, w, r)
        } else {
            writeError(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server running")
    log.Fatal(http.ListenAndServe(":8080", logHandler(mux)))
}
