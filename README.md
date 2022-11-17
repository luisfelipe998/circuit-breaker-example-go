# circuit breaker example in GO

This is a simple client -> server interaction of a circuit breaker. The objective of circuit breakers are to protect remote calls to system that eventually will be unavailable for some time.
The repo has two examples: one using `gobreaker` and another one with `own implementation` using [state design pattern](https://refactoring.guru/design-patterns/state).

## gobreaker

Go to `client/server` and comment line 8 and uncomment line 9. 
This enables the implementation of the circuit breaker using the `github.com/sony/gobreaker` lib through a simple [decorator pattern](https://refactoring.guru/design-patterns/decorator).
To run the project, make sure you have `docker` installed and run `make run`

## own implementation

Go to `client/server` and comment line 8 and uncomment line 10. 
This enables the implementation of the circuit breaker using the my own simple implementation under `internal/circuitbreaker` folder through a simple [decorator pattern](https://refactoring.guru/design-patterns/decorator).
To run the project, make sure you have `docker` installed and run `make run`