package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func doWork(sec int, channel chan string) {
	time.Sleep(time.Duration(sec) * time.Second)
	channel <- fmt.Sprintf("Work: %d times", rand.IntN(100))
	wg.Done()
}


var wg sync.WaitGroup

func main() {
	start := time.Now()

	channel := make(chan string)

	wg.Add(3)
	go doWork(1, channel)
	go doWork(2, channel)
	go doWork(3, channel)
	
	go func() {
		for res := range channel {
			fmt.Println(res)
		}
	}()

	wg.Wait()
	close(channel)
	fmt.Println("Main goroutine takes ", time.Since(start))
}

