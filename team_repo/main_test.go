package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestCreateOrg(t *testing.T) {
    orgData := map[string]string{"name": "testOrg"}

    jsonData, err := json.Marshal(orgData)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := http.Post("http://localhost:3000/createOrg", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected Status OK, got %v", resp.Status)
    }

    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }

    var responseOrgData map[string]interface{}
    if err := json.Unmarshal(responseBody, &responseOrgData); err != nil {
        t.Fatal(err)
    }

    if responseOrgData["name"] != "testOrg" {
        t.Errorf("Expected org name to be 'testOrg', got '%v'", responseOrgData["name"])
    }
}
