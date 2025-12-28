package basics

import (
	"fmt"
	"sync"
	// "time"
)

func Task (number int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Number", number)
}


func Routine () {
	var waitgroup sync.WaitGroup
	for i := range 10{
		waitgroup.Add(1)
		go Task(i, &waitgroup)
	}

	waitgroup.Wait()
}
