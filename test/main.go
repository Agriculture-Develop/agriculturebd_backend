package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type MyConcurrentMap struct {
	m map[int]int
	sync.RWMutex
	c chan struct{}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
	m.c <- struct{}{}
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	defer m.Unlock()

	for {
		if v, ok := m.m[k]; ok {
			return v, nil
		}
		select {
		case <-time.After(maxWaitingDuration):
			return 0, errors.New("timeout")
		case <-m.c:
			if v, ok := m.m[k]; ok {
				return v, nil
			}
		}
	}
}

func main() {
	myMap := &MyConcurrentMap{
		c: make(chan struct{}, 1),
		m: make(map[int]int),
	}

	go func() {
		fmt.Println(myMap.Get(1, 3*time.Second))
	}()

	myMap.Put(1, 2)
	time.Sleep(4 * time.Second)
}
