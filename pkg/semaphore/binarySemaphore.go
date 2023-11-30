package semaphore

type BinarySemaphore struct {
	state bool
	queue chan chan struct{}
}

func NewBinarySemaphore() *BinarySemaphore {
	return &BinarySemaphore{
		state: true,
		queue: make(chan chan struct{}),
	}
}

func (bs *BinarySemaphore) Wait() {
	waitSignal := make(chan struct{})

	if bs.state {
		bs.state = false
	} else {
		bs.queue <- waitSignal
		<-waitSignal
	}
}

func (bs *BinarySemaphore) Signal() {
	select {
	case waitSignal := <-bs.queue:
		close(waitSignal)
	default:
		bs.state = true
	}
}
