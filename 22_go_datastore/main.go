package main

import (
	"fmt"
	"time"
)

// type DataStore interface {
// 	Set() (bool, error)
// 	Get() (string, error)
// 	Delete() (bool, error)
// }

type DataStore struct {
	data []KeyValuePair
}

type KeyValuePair struct {
	key   string
	value string
	ttl   time.Duration
}

func (s *DataStore) Get(key string) (string, error) {

	var kv *KeyValuePair
	for _, KV := range s.data {
		if KV.key == key {
			kv = &KV
		}
	}

	if kv == nil {
		return "", fmt.Errorf("not found")
	}

	return kv.value, nil
}

func (s *DataStore) Set(key string, val string, ttl time.Duration) (bool, error) {

	var kv *KeyValuePair
	for _, KV := range s.data {
		if KV.key == key {
			kv = &KV
		}
	}

	if kv == nil {
		s.data = append(s.data, KeyValuePair{key: key, value: val, ttl: ttl})

		go func() {
			time.Sleep(ttl)
			fmt.Printf("In sleep")
		}()
		return true, nil
	}

	kv.value = val
	kv.ttl = ttl
	return true, nil
}

func (s *DataStore) Delete(key string) (bool, error) {
	for i, KV := range s.data {
		if KV.key == key {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("key not found")
}

func main() {
	ds := DataStore{data: []KeyValuePair{}}

	ds.Set("key1", "value1", 500)

	a, _ := ds.Get("key1")

	fmt.Println(a)

	// var ds DataStore = {data: []KeyValuePair}

}
