package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)

func TestPutHandler(t *testing.T) {
    kv := &KeyValueStore{}
    req := httptest.NewRequest("PUT", "/put/testkey", strings.NewReader("testvalue"))
    w := httptest.NewRecorder()

    putHandler(kv, "testkey", w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }
}

func TestGetHandler(t *testing.T) {
    kv := &KeyValueStore{}
    kv.Put("testkey", "testvalue")

    req := httptest.NewRequest("GET", "/get/testkey", nil)
    w := httptest.NewRecorder()

    getHandler(kv, "testkey", w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }
}

func TestDeleteHandler(t *testing.T) {
    kv := &KeyValueStore{}
    kv.Put("testkey", "testvalue")

    req := httptest.NewRequest("DELETE", "/delete/testkey", nil)
    w := httptest.NewRecorder()

    deleteHandler(kv, "testkey", w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }
}

func TestListKeysHandler(t *testing.T) {
    kv := &KeyValueStore{}
    kv.Put("testkey1", "value1")
    kv.Put("testkey2", "value2")

    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    listKeysHandler(kv, w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }
}

func TestParallelism(t *testing.T) {
    kv := &KeyValueStore{}

    t.Run("Parallel", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            go func(i int) {
                kv.Put("key", "value")
                kv.Get("key")
                kv.Delete("key")
            }(i)
        }
    })

    t.Run("Parallel with different keys", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            go func(i int) {
                key := "key" + strconv.Itoa(i)
                kv.Put(key, "value")
                kv.Get(key)
                kv.Delete(key)
            }(i)
        }
    })
}
