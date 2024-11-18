package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tianrosandhy/goadaptivethread"
)

// main demonstrates how to use AdaptiveThread to execute 100000 function calls concurrently.
//
// It showcases that even though we attempt to execute 100000 goroutines, AdaptiveThread
// limits it to 100 goroutines only, making it safe to use on a system with limited resources.

func main() {
	start := time.Now().UnixNano()

	at := goadaptivethread.NewAdaptiveThread(100) // will spawn maximum 100 goroutines only

	for i := 0; i < 100000; i++ {
		at.Execute(func() {
			worker(i) // will be run asynchronously
		})
	}

	at.WaitComplete() // will wait until all executed

	fmt.Println("--------------")
	fmt.Printf("Finish in %.3fs", float64(time.Now().UnixNano()-start)/1e9)
}

func worker(iteration int) {
	log.Printf("[WORKER] ITERATION %d", iteration)
}
