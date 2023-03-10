package integration

import (
	"encoding/json"
	"fmt"
	"github.com/Eitol/newsletter-api/pkg/newsletter/delivery/httpserver"
	"net/http"
	"strings"
	"testing"
)

var testPort = 8081

func TestSubscriptionsPost(t *testing.T) {
	go httpserver.RunHttpServer(testPort)
	awaitForServer(testPort)
	body := strings.NewReader(`{
    		"userId": "b3a5fcab-4ef2-402b-99d5-6beb882bcf81",
    		"blogId": "b3a5fcab-4ef2-402b-99d5-6beb882bcf8b",
    		"interests": ["tech", "sport"]
	}`)
	url := fmt.Sprintf("http://localhost:%d/newsletter/subscriptions", testPort)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}
	respMap := make(map[string]string)
	err = json.NewDecoder(resp.Body).Decode(&respMap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	sId, ok := respMap["subscriptionId"]
	if !ok {
		t.Errorf("Expected subscriptionId in response body but got %v", respMap)
	}
	if len(sId) != 36 {
		t.Errorf("Expected subscriptionId to be 36 characters long but got %d", len(sId))
	}
}
