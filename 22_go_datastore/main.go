package main

import (
	"fmt"
	"sync"
	"time"
)

// type DataStore interface {
// 	Set() (bool, error)
// 	Get() (string, error)
// 	Delete() (bool, error)
// }

type DataStore struct {
	data map[string]Data
}

type Data struct {
	value string
	ttl   time.Duration
}

func (s *DataStore) Get(key string) (string, error) {
	data, exists := s.data[key]

	if !exists {
		return "", fmt.Errorf("not found")
	}

	return data.value, nil
}

func (s *DataStore) Set(key string, val string, ttl time.Duration, wg *sync.WaitGroup) (bool, error) {

	data := Data{value: val, ttl: ttl}
	s.data[key] = data

	go func() {
		defer wg.Done()
		fmt.Printf("sleeping now!\n")
		time.Sleep(ttl * time.Second)
		fmt.Printf("after sleep deleting now!\n")
		s.Delete(key)
	}()

	return true, nil
}

func (s *DataStore) Delete(key string) (bool, error) {

	_, exists := s.data[key]

	if !exists {
		return false, fmt.Errorf("not found")
	}

	delete(s.data, key)
	return true, nil
}

func main() {
	ds := DataStore{data: make(map[string]Data)}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	ds.Set("key1", "value1", 5, wg)

	time.Sleep(3 * time.Second)
	fmt.Printf("woke up 1 now!\n")
	a, _ := ds.Get("key1")
	fmt.Println(a)

	time.Sleep(3 * time.Second)
	fmt.Printf("woke up 2 now!\n")
	a, err := ds.Get("key1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)

	wg.Wait()

}
