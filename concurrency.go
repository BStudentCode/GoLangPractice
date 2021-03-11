package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	//add call to fmt.Scanln() to prevent program instantly closing
	wg.Wait() // waits until goroutines are finished then terminates
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
