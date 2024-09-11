package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func putHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid body", http.StatusBadRequest)
        return
    }

    kv.Put(key, string(body))
    fmt.Fprintf(w, "Stored value for key: %s", key)
}

func getHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    value, ok := kv.Get(key)
    if !ok {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Value for key %s: %s", key, value)
}

func deleteHandler(kv *KeyValueStore, key string, w http.ResponseWriter, r *http.Request) {
    kv.Delete(key)
    fmt.Fprintf(w, "Deleted key: %s", key)
}

func listKeysHandler(kv *KeyValueStore, w http.ResponseWriter, r *http.Request) {
    keys := kv.ListKeys()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(keys)
}
