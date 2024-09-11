package main

import (
    "sync"
)

type KeyValueStore struct {
    store sync.Map
}

func (kv *KeyValueStore) Put(key string, value string) {
    kv.store.Store(key, value)
}

func (kv *KeyValueStore) Get(key string) (string, error) {
    value, ok := kv.store.Load(key)
    if ok {
        return value.(string), nil
    }
    return "", ErrNotFound
}

func (kv *KeyValueStore) Delete(key string) error {
    _, ok := kv.store.Load(key)
    if !ok {
        return ErrNotFound
    }
    kv.store.Delete(key)
    return nil
}

func (kv *KeyValueStore) ListKeys() []string {
    var keys []string
    kv.store.Range(func(key, value interface{}) bool {
        keys = append(keys, key.(string))
        return true
    })
    return keys
}
