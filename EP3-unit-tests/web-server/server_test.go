package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a new request with GET method and no request body
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler with the request and the recorder
	handler(rr, req)

	// Check the status code of the response
	gotStatus := rr.Code
	wantStatus := http.StatusOK
	if gotStatus != wantStatus {
		t.Errorf("handler returned wrong status code: got %v, want %v", gotStatus, wantStatus)
	}

	// Check the body of the response
	gotBody := rr.Body.String()
	wantBody := "Hello, world!"
	if gotBody != wantBody {
		t.Errorf("handler returned unexpected body: got %v, want %v", gotBody, wantBody)
	}
}
