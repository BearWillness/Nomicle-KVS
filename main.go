package main

import (
    "log"
    "net/http"
)

func main() {
    kvStore := &KeyValueStore{}

    http.HandleFunc("/put/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Path[len("/put/"):]
        if r.Method == "PUT" {
            putHandler(kvStore, key, w, r)
        } else {
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Path[len("/get/"):]
        if r.Method == "GET" {
            getHandler(kvStore, key, w, r)
        } else {
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Path[len("/delete/"):]
        if r.Method == "DELETE" {
            deleteHandler(kvStore, key, w, r)
        } else {
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            listKeysHandler(kvStore, w, r)
        } else {
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server running")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
