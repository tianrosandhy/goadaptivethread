package goadaptivethread

type AdaptiveThread struct {
	WorkerFunc func()

	threadCount int
	threadChan  chan bool
}

// NewAdaptiveThread will initiate a new AdaptiveThread struct with its used thread count
func NewAdaptiveThread(threadCount int) *AdaptiveThread {
	a := AdaptiveThread{
		threadCount: threadCount,
		threadChan:  make(chan bool, threadCount),
	}

	return &a
}

// Execute will run workerFunc() asynchronously, with maximum of threadCount goroutine spawned
func (a *AdaptiveThread) Execute(workerFunc func()) {
	a.threadChan <- true
	go func() {
		defer func() {
			<-a.threadChan
		}()
		workerFunc()
	}()

	// how to call a.WorkerFunc dynamically
}

// WaitComplete will wait until all execution is finished
func (a *AdaptiveThread) WaitComplete() {
	for i := 0; i < cap(a.threadChan); i++ {
		a.threadChan <- true
	}
	close(a.threadChan)
}
