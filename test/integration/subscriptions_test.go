package integration

import (
	"encoding/json"
	"fmt"
	"github.com/Eitol/newsletter-api/pkg/newsletter/delivery/httpserver"
	"github.com/Eitol/newsletter-api/pkg/newsletter/handler"
	"net/http"
	"strings"
	"testing"
)

var testPort = 8081

func TestSubscriptions(t *testing.T) {
	go httpserver.RunHttpServer(testPort)
	awaitForServer(testPort)
	body := strings.NewReader(`{
    		"userId": "13a5fcab-4ef2-402b-99d5-6beb882bcf81",
    		"blogId": "13a5fcab-4ef2-402b-99d5-6beb882bcf81",
    		"interests": ["tech", "sport"]
	}`)
	id1 := postNewSubscription(t, body)

	body = strings.NewReader(`{
    		"userId": "23a5fcab-4ef2-402b-99d5-6beb882bcf82",
    		"blogId": "23a5fcab-4ef2-402b-99d5-6beb882bcf82",
    		"interests": ["politics"]
	}`)
	id2 := postNewSubscription(t, body)

	if id1 == id2 {
		t.Errorf("Expected different ids but got %s and %s", id1, id2)
	}

	all := getSubscription(t, "", "", []string{})
	if len(all.Results) != 2 {
		t.Errorf("Expected 2 subscriptions but got %v", all)
	}
	s1 := getSubscription(t, "13a5fcab-4ef2-402b-99d5-6beb882bcf81", "", []string{})
	if len(s1.Results) != 1 {
		t.Errorf("Expected 1 subscription but got %v", s1)
	}
	if s1.Results[0].UserID != "13a5fcab-4ef2-402b-99d5-6beb882bcf81" {
		t.Errorf("Expected user id 13a5fcab-4ef2-402b-99d5-6beb882bcf81 but got %s", s1.Results[0].UserID)
	}
	s2 := getSubscription(t, "", "13a5fcab-4ef2-402b-99d5-6beb882bcf81", []string{})
	if len(s2.Results) != 1 {
		t.Errorf("Expected 1 subscription but got %v", s2)
	}
	if s2.Results[0].BlogID != "13a5fcab-4ef2-402b-99d5-6beb882bcf81" {
		t.Errorf("Expected blog id 13a5fcab-4ef2-402b-99d5-6beb882bcf81 but got %s", s2.Results[0].BlogID)
	}

}

func getSubscription(t *testing.T, userID string, blogID string, interest []string) handler.ResponseDoc {
	url := fmt.Sprintf("http://localhost:%d/newsletter/subscriptions?page=1&maxPageSize=10", testPort)
	url = addParamsToUrl(url, userID, blogID, interest)
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}
	var respObjArr []handler.ResponseDoc
	err = json.NewDecoder(resp.Body).Decode(&respObjArr)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(respObjArr) == 0 {
		return handler.ResponseDoc{}
	}
	respObj := respObjArr[0]
	if len(respObj.Results) == 0 {
		t.Errorf("Expected at least one subscription but got %v", respObj)
	}
	for _, sub := range respObj.Results {
		if userID != "" && sub.UserID != userID {
			t.Errorf("Expected userID %s but got %s", userID, sub.UserID)
		}
		if blogID != "" && sub.BlogID != blogID {
			t.Errorf("Expected blogID %s but got %s", blogID, sub.BlogID)
		}
		if len(interest) > 0 {
			for _, i := range interest {
				for _, j := range sub.Interests {
					if i != j {
						t.Errorf("Expected interest %s but got %s", i, j)
					}
				}
			}
		}
	}
	return respObj
}

func addParamsToUrl(url, userID, blogID string, interest []string) string {
	if userID != "" {
		url += "&userId=" + userID
	}
	if blogID != "" {
		url += "&blogId=" + blogID
	}
	if len(interest) > 0 {
		url += "&interest=" + interest[0]
		for i := 1; i < len(interest); i++ {
			url += "&interest=" + interest[i]
		}
	}
	return url
}

func postNewSubscription(t *testing.T, body *strings.Reader) string {
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
	return sId
}
