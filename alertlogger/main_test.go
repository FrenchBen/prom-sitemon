package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// Add a handler for the function we want to test
	hdl := http.HandlerFunc(handler)

	// Create a table driven test payload
	var payloadTests = []struct {
		name     string
		body     string
		status   int
		expected string
	}{
		{"ValidRequest", `{"receiver": "test", "alerts": [{}]}`, http.StatusOK, "OK"},
		{"EmptyBody", ``, http.StatusBadRequest, "Please send a request body"},
		{"BadJson", `{"bad":json}`, http.StatusBadRequest, "invalid character"},
		{"NoAlerts", `{"receiver": "test"}`, http.StatusBadRequest, "No alerts to display"},
	}

	for _, payloadTest := range payloadTests {
		t.Run(payloadTest.name, func(t *testing.T) {
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(payloadTest.body))
			// Send the request using the request and response recorder created
			hdl.ServeHTTP(rr, req)
			// Check the status code is what we expect.
			// assert.Equal(t, payloadTest.status, rr.Code)
			if status := rr.Code; status != payloadTest.status {
				t.Errorf("handler returned wrong status code: got `%v` want `%v`", status, payloadTest.status)
			}
			// Check the response body is what we expect
			actual := strings.TrimSpace(rr.Body.String())
			if actual != payloadTest.expected && !strings.HasPrefix(actual, payloadTest.expected) {
				t.Errorf("handler returned unexpected body: got `%v` want `%v`", actual, payloadTest.expected)
			}
		})
	}

}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("GET", "/_health", nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	hdl := http.HandlerFunc(healthCheckHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	hdl.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
