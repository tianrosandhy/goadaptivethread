# Goadaptivethread
The goadaptivethread package provides a mechanism for limiting the number of concurrent goroutines executing a task, ensuring that only a specified maximum number of goroutines are active at any given time. This is useful for applications where you want to control concurrency, such as managing system resources or working with rate-limited APIs.

## Features
1. **Concurrency Control**: The AdaptiveThread struct leverages a buffered channel to restrict the number of concurrent goroutines to threadCount. Tasks block on the channel when the maximum concurrency is reached, resuming only when a running goroutine completes.

2. **Flexible Task Execution**: `Execute(workerFunc func())` accepts a func() as the worker function. This design supports closures, allowing you to dynamically capture variables and pass arguments into tasks. This means you can use any function, including those with different arguments, by wrapping them in a closure.

3. **Graceful Shutdown with WaitClose()**: The WaitClose() method waits for all goroutines to complete before the program exits, helping to prevent prematurely ending tasks and ensuring a controlled shutdown process.


## Use Cases
1. Resource Management: Use adaptivethread to control memory and CPU usage by limiting concurrent tasks. (Too many active uncontrolled goroutine possibly bring uncontrollable memory leak)
2. API Rate Limiting: Maintain the rate limit when making network requests by setting a maximum number of concurrent requests.
3. Batch Processing: Process large datasets concurrently with control over system resources. 


## Usage Example
```
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
```