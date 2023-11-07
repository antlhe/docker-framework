package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
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
    defer resp.Body.Close() // Ensure we close the body

    // Check the response
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected Status OK, got %v", resp.Status)
    }

    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }

    // Unmarshal the JSON data into a map
    var responseOrgData map[string]interface{}
    if err := json.Unmarshal(responseBody, &responseOrgData); err != nil {
        t.Fatal(err)
    }

    // Assert that the 'name' field in the response matches 'testOrg'
    if responseOrgData["name"] != "testOrg" {
        t.Errorf("Expected org name to be 'testOrg', got '%v'", responseOrgData["name"])
    }
}
