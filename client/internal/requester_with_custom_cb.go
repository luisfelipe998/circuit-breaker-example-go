package internal

import "github.tools.sap/I521862/circuit-breaker-demo/client/internal/circuitbreaker"

func NewCustomCircuitBreakerRequester() CustomCircuitBreakerRequester {
	return CustomCircuitBreakerRequester{
		defaultRequester: NewDefaultRequester(),
		circuitBreaker:   circuitbreaker.NewCircuitBreaker(),
	}
}

type CustomCircuitBreakerRequester struct {
	defaultRequester Requester
	circuitBreaker   *circuitbreaker.CircuitBreaker
}

func (r CustomCircuitBreakerRequester) MakeGetRequest(endpoint string) (ServerResponse, error) {
	body, err := r.circuitBreaker.Protect(func() (any, error) {
		return r.defaultRequester.MakeGetRequest(endpoint)
	})
	if err != nil {
		return ServerResponse{}, err
	}

	return body.(ServerResponse), nil
}
