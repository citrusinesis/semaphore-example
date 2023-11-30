package semaphore

import "sync"

type BinarySemaphore struct {
	state bool
	queue chan chan struct{}
	mu    sync.Mutex
}

func NewBinarySemaphore() *BinarySemaphore {
	return &BinarySemaphore{
		state: true,
		queue: make(chan chan struct{}),
	}
}

func (bs *BinarySemaphore) Wait() {
	waitSignal := make(chan struct{})

	bs.mu.Lock()
	if bs.state {
		bs.state = false
		bs.mu.Unlock()
	} else {
		bs.queue <- waitSignal
		bs.mu.Unlock()
		<-waitSignal
	}
}

func (bs *BinarySemaphore) Signal() {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	select {
	case waitSignal := <-bs.queue:
		close(waitSignal)
	default:
		bs.state = true
	}
}
