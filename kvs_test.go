package main

import (
    "sync"
    "testing"
    "fmt"
)

func TestKeyValueStoreParallel(t *testing.T) {
    kv := KeyValueStore{}
    wg := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            kv.Put("key", "value")
            _, _ = kv.Get("key")
            kv.Delete("key")
        }(i)
    }

    wg.Wait()
}

func TestKeyValueStoreParallelDifferentKeys(t *testing.T) {
    kv := KeyValueStore{}
    wg := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", i)
            kv.Put(key, fmt.Sprintf("value%d", i))
            _, _ = kv.Get(key)
            kv.Delete(key)
        }(i)
    }

    wg.Wait()
}
