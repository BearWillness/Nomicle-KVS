package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "errors"
)

var ErrNotFound = errors.New("key not found")

func putHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        http.Error(w, "Missing key", http.StatusBadRequest)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid body", http.StatusBadRequest)
        return
    }

    kv.Put(key, string(body))
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "success", "key": key})
}

func getHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        http.Error(w, "Missing key", http.StatusBadRequest)
        return
    }

    value, ok := kv.Get(key)
    if !ok {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
}

func deleteHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        http.Error(w, "Missing key", http.StatusBadRequest)
        return
    }

    _, ok := kv.Get(key)
    if !ok {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }

    kv.Delete(key)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "deleted", "key": key})
}

func listKeysHandler(kv *KeyValueStore, w http.ResponseWriter, r *http.Request) {
    keys := kv.ListKeys()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(keys)
}
