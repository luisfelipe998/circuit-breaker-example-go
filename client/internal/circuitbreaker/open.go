package circuitbreaker

import (
	"fmt"
	"time"
)

var (
	ErrCircuitBreakerOpen = fmt.Errorf("circuit breaker still open")
)

type Open struct {
	circuitBreaker *CircuitBreaker
}

func NewOpen(cb *CircuitBreaker) State {
	return &Open{
		circuitBreaker: cb,
	}
}

func (o *Open) Execute(callback func() (any, error)) (any, error) {
	if time.Now().After(o.circuitBreaker.resetTimeAfterTrip) {
		o.circuitBreaker.prepareReset()
		return o.circuitBreaker.currentState.Execute(callback)
	} else {
		fmt.Println(ErrCircuitBreakerOpen.Error())
		return nil, ErrCircuitBreakerOpen
	}
}
