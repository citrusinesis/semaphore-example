package semaphore

type Barrier struct {
	bCount   int
	bTotal   int
	sbarrier *Semaphore
	bmutex   *Semaphore
}

func NewBarrier(n int) *Barrier {
	return &Barrier{
		bCount:   0,
		bTotal:   n,
		sbarrier: NewSemaphore(0),
		bmutex:   NewSemaphore(1),
	}
}

func (b *Barrier) BarrierWait() {
	b.bmutex.Wait()
	b.bCount++
	b.bmutex.Signal()

	if b.bCount == b.bTotal {
		b.sbarrier.Signal()
	}

	b.sbarrier.Wait()
	b.sbarrier.Signal()
}
