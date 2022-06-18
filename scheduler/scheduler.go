package scheduler

import (
	"fmt"
	"time"
)

// Scheduler will run a
type Scheduler[T any] struct {
	afterFunc schFunc[T]
	close     chan struct{}
	// Error is a field to allow assignment
	// in chain, and just check the Error at the end.
	Error     error
	isRunning bool
	mainFunc  schFunc[T]
	period    time.Duration
	preFunc   schFunc[T]
	value     T
}

type schFunc[T any] func(i *T) error

// NewScheduler returns a new instance of scheduler.
func NewScheduler[T any](period time.Duration, value T) *Scheduler[T] {
	return &Scheduler[T]{
		value:  value,
		period: period,
		close:  make(chan struct{}, 1),
	}
}

// Close will stop the Run function, and run the AfterFunc.
func (s *Scheduler[T]) Close() error {
	if !s.isRunning {
		return fmt.Errorf("the scheduler is not running")
	}
	s.close <- struct{}{}
	return nil
}

// Run will start the process, and only stop when Close() is called.
func (s *Scheduler[T]) Run() error {
	ticker := time.NewTicker(s.period)

	if s.preFunc != nil {
		s.preFunc(&s.value)
	}

	for {
		select {
		case <-s.close:
			if s.afterFunc != nil {
				return s.afterFunc(&s.value)
			}
			return nil
		case <-ticker.C:
			s.mainFunc(&s.value)
		}
	}
}

// SetAfterFunc will set the function that will be ran
// just once after the MainFunc.
func (s *Scheduler[T]) SetAfterFunc(f schFunc[T]) *Scheduler[T] {
	if s.Error != nil {
		return s
	}
	if f == nil {
		s.Error = fmt.Errorf("the function 'SetAfterFunc' received a nil function")
		return s
	}
	s.afterFunc = f
	return s
}

// SetPreFunc will set the function that will be ran
// just once before .
func (s *Scheduler[T]) SetPreFunc(f schFunc[T]) *Scheduler[T] {
	if s.Error != nil {
		return s
	}
	if f == nil {
		s.Error = fmt.Errorf("the function 'SetPreFunc' received a nil function")
		return s
	}
	s.preFunc = f

	return s
}

// SetMain will set the function that will be ran
// in a period of time.
// *required.
func (s *Scheduler[T]) SetMain(f schFunc[T]) *Scheduler[T] {
	if s.Error != nil {
		return s
	}
	if f == nil {
		s.Error = fmt.Errorf("the function 'SetMain' received a nil function")
		return s
	}
	s.mainFunc = f

	return s
}
