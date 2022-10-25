package internal

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreakerRequester() CircuitBreakerRequester {
	circuitBreakerConfig := gobreaker.Settings{
		Name:        "/server/test",
		MaxRequests: 3,
		Interval:    time.Second * 10,
		Timeout:     time.Second * 20,
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("circuit breaker %s changed state from %s to %s", name, from, to)
		},
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			log.Printf("consecutive server failures: %d", counts.ConsecutiveFailures)
			return counts.ConsecutiveFailures > 2
		},
	}

	return CircuitBreakerRequester{
		defaultRequester: NewDefaultRequester(),
		circuitBreaker:   gobreaker.NewCircuitBreaker(circuitBreakerConfig),
	}
}

type CircuitBreakerRequester struct {
	defaultRequester Requester
	circuitBreaker   *gobreaker.CircuitBreaker
}

func (r CircuitBreakerRequester) MakeGetRequest(endpoint string) (ServerResponse, error) {
	body, err := r.circuitBreaker.Execute(func() (interface{}, error) {
		return r.defaultRequester.MakeGetRequest(endpoint)
	})
	if err != nil {
		return ServerResponse{}, err
	}

	return body.(ServerResponse), nil
}
