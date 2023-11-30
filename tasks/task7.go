package tasks

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type SyncMap struct {
	mtx sync.RWMutex
	data map[string]int
}

func (s *SyncMap) Add(add_string string, data int) error{
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, found := s.data[add_string]
	if found {
		return errors.New("data already in map")
	}
	s.data[add_string] = data
	return nil
}

func (s *SyncMap) Get(g string) (int, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	ans, found := s.data[g]
	if found {
		return ans, nil
	}
	return 0, errors.New("no data in map")
}


func Test7 () bool {
	base := make([]int, 100)

	wg := sync.WaitGroup{}

	sm := &SyncMap{
		data: map[string]int{},
	}

	for i := 0; i < len(base); i++{
		base[i] = rand.Intn(50)
	}

	for idx, num := range base{
		wg.Add(1)
		go func(n, num int){
			defer wg.Done()
			err := sm.Add(strconv.Itoa(n), num)
			if err != nil{
				log.Fatalf("Failed to add key %v\n", err)
			}
		}(idx, num)
	}

	for idx, _ := range base{
		wg.Add(1)
		go func(n int){
			defer wg.Done()
			n, err := sm.Get(strconv.Itoa(n))
			if err != nil{
				log.Fatalf("Failed to add key %v\n", err)
				return
			}
			fmt.Println(n)
		}(idx)
	}
	wg.Wait()
	return true
}