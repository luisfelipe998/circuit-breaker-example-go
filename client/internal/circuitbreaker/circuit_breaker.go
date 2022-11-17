package circuitbreaker

import (
	"fmt"
	"sync"
	"time"
)

type State interface {
	Execute(func() (any, error)) (any, error)
}

type CircuitBreaker struct {
	closed   State
	halfOpen State
	open     State

	currentState       State
	failureCount       int
	successCount       int
	tripTimeout        time.Duration
	resetTimeAfterTrip time.Time
	mutex              *sync.Mutex
}

func NewCircuitBreaker() *CircuitBreaker {
	cb := &CircuitBreaker{
		failureCount: 0,
		successCount: 0,
		tripTimeout:  time.Second * 15,
		mutex:        &sync.Mutex{},
	}

	cb.closed = NewClosed(cb)
	cb.halfOpen = NewHalfOpen(cb)
	cb.open = NewOpen(cb)

	cb.reset()
	return cb
}

func (cb *CircuitBreaker) trip() {
	fmt.Println("circuit breaker switched to open state")
	cb.setState(cb.open)
}

func (cb *CircuitBreaker) prepareReset() {
	fmt.Println("circuit breaker switched to half-open state")
	cb.setState(cb.halfOpen)
}

func (cb *CircuitBreaker) reset() {
	fmt.Println("circuit breaker switched to closed state")
	cb.setState(cb.closed)
}

func (cb *CircuitBreaker) setState(s State) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.failureCount = 0
	cb.successCount = 0
	cb.resetTimeAfterTrip = time.Now().Add(cb.tripTimeout)
	cb.currentState = s
}

func (cb *CircuitBreaker) incrementFailureCount() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.failureCount++
	fmt.Printf("failureCount: %d\n", cb.failureCount)
}

func (cb *CircuitBreaker) incrementSuccessCount() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.successCount++
	fmt.Printf("successCount: %d\n", cb.successCount)
}

func (cb *CircuitBreaker) Protect(callback func() (any, error)) (any, error) {
	return cb.currentState.Execute(callback)
}
