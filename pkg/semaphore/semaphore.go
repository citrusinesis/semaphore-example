package semaphore

type Semaphore struct {
	count int
	queue chan chan struct{}
}

func NewSemaphore(count int) *Semaphore {
	return &Semaphore{
		count: count,
		queue: make(chan chan struct{}),
	}
}

func (s *Semaphore) Wait() {
	waitSignal := make(chan struct{})
	s.count--

	if s.count < 0 {
		s.queue <- waitSignal
		<-waitSignal
	} else {
		close(waitSignal)
	}
}

func (s *Semaphore) Signal() {
	s.count++

	if s.count <= 0 {
		waitSignal := <-s.queue
		close(waitSignal)
	}
}
