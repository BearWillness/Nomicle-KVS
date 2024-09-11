package main

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
)

var ErrNotFound = errors.New("key not found")

func writeError(w http.ResponseWriter, errMsg string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{"error": errMsg})
}

func putHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        writeError(w, "Missing key", http.StatusBadRequest)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        writeError(w, "Invalid body", http.StatusBadRequest)
        return
    }

    kv.Put(key, string(body))
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "success", "key": key})
}

func getHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        writeError(w, "Missing key", http.StatusBadRequest)
        return
    }

    value, err := kv.Get(key)
    if errors.Is(err, ErrNotFound) {
        writeError(w, "Not found", http.StatusNotFound)
        return
    } else if err != nil {
        writeError(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
}

func deleteHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    if key == "" {
        writeError(w, "Missing key", http.StatusBadRequest)
        return
    }

    err := kv.Delete(key)
    if errors.Is(err, ErrNotFound) {
        writeError(w, "Not found", http.StatusNotFound)
        return
    } else if err != nil {
        writeError(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "deleted", "key": key})
}

func listKeysHandler(kv *KeyValueStore, w http.ResponseWriter, r *http.Request) {
    keys := kv.ListKeys()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(keys)
}
