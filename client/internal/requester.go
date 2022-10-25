package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ServerResponse struct {
	Ok bool `json:"ok"`
}

func NewDefaultRequester() DefaultRequester {
	return DefaultRequester{}
}

type DefaultRequester struct{}

func (r DefaultRequester) MakeGetRequest(endpoint string) (ServerResponse, error) {
	serverEndpoint := fmt.Sprintf("http://%s:%s%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"), endpoint)
	log.Printf("calling server: %s", serverEndpoint)

	resp, err := http.Get(serverEndpoint)
	if err != nil {
		return ServerResponse{}, fmt.Errorf("failed to perform get request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return ServerResponse{}, fmt.Errorf("server failed to process the request. Returned internal server error")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServerResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	var responseBody ServerResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return ServerResponse{}, fmt.Errorf("failed to parse response body: %w", err)
	}
	return responseBody, nil
}
