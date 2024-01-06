package go_sema

type Semaphore interface {
	Acquire()
	Release()
	Len() int
	IsEmpty() bool
}

type semaphore struct {
	semC chan struct{}
}

func safe(maxConcurrency int) int {
	if maxConcurrency == -1 {
		maxConcurrency = 333_333
	}

	if maxConcurrency < 1 {
		maxConcurrency = 1
	}

	return maxConcurrency
}

func New(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, safe(maxConcurrency)),
	}
}

func (s *semaphore) IsEmpty() bool {
	return s.Len() == 0
}

func (s *semaphore) Len() int {
	return len(s.semC)
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}
