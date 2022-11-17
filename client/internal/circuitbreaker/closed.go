package circuitbreaker

type Closed struct {
	circuitBreaker *CircuitBreaker
}

func NewClosed(cb *CircuitBreaker) State {
	return &Closed{
		circuitBreaker: cb,
	}
}

func (c *Closed) Execute(callback func() (any, error)) (any, error) {
	response, err := callback()
	if err != nil {
		c.circuitBreaker.incrementFailureCount()
	}

	if c.circuitBreaker.failureCount > 2 {
		c.circuitBreaker.trip()
	}

	return response, err
}
