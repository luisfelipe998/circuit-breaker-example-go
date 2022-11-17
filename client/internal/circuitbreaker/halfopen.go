package circuitbreaker

type HalfOpen struct {
	circuitBreaker *CircuitBreaker
}

func NewHalfOpen(cb *CircuitBreaker) State {
	return &HalfOpen{
		circuitBreaker: cb,
	}
}

func (h *HalfOpen) Execute(callback func() (any, error)) (any, error) {
	response, err := callback()
	if err != nil {
		h.circuitBreaker.incrementFailureCount()
	} else {
		h.circuitBreaker.incrementSuccessCount()
	}

	if h.circuitBreaker.failureCount > 0 {
		h.circuitBreaker.trip()
	} else if h.circuitBreaker.successCount > 1 {
		h.circuitBreaker.reset()
	}

	return response, err
}
