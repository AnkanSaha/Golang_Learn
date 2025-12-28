package basics

import (
	"fmt"
	"sync"
)

type value struct {
	values int
	mu sync.Mutex
}

func (va *value) inc(wg *sync.WaitGroup) {
	defer func () {wg.Done()}()
	va.mu.Lock()
	va.values += 1
	va.mu.Unlock()
}

func StartMutex() {
	var newVar = value{values: 0}
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)
		go newVar.inc(&wg)
	}
	wg.Wait()
	fmt.Println(newVar.values)
}