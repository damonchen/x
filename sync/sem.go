package sync

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(n int) Semaphore {
	return Semaphore{
		ch: make(chan struct{}, n),
	}
}

func (s Semaphore) Accquire() {
	select {
	case s.ch <- struct{}{}:
		return
	}
}

func (s Semaphore) Release() {
	<-s.ch
}
