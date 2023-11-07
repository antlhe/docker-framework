package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateOrg(t *testing.T) {
	// Sample test data
	orgData := map[string]string{"name": "testOrg"}

	// Convert to JSON
	jsonData, err := json.Marshal(orgData)
	if err != nil {
		t.Fatal(err)
	}

	// Call the mock createOrg function
	resp, err := http.Post("http://localhost:3000/createOrg", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected Status OK, got %v", resp.Status)
	}

	// TODO: Add more assertions to validate response content
}
