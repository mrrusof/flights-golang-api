package main

import (
  "log"
  "net/http"
  "testing"
  "os"
  "fmt"
  "strings"
)

const HOST = "localhost"
const PORT = 8080

func TestMain(t *testing.T) {
  if os.Getenv("E2E") != "true" {
    t.Skip("Skipping E2E tests.")
  }

  cases := []struct {
    requestBody string
    expectedStatusCode int
  }{
    {"not a JSON doc", 400},
    {"{}", 400},
    {"{\"path\": [{\"origin\": \"ATL\", \"destination\": \"EWR\"}, {\"origin\": \"SFO\", \"destination\": \"ATL\"}]}", 200},
  }

  url := fmt.Sprintf("http://%s:%d/flights/summarize", HOST, PORT)

  for _, c := range cases {
    resp, err := http.Post(url, "application/json", strings.NewReader(c.requestBody))

    if err != nil {
      log.Fatalf("Failed to POST: %s", err)
    }

    if resp.StatusCode != c.expectedStatusCode {
      t.Errorf("Wrong response status: expected %d, got %d", c.expectedStatusCode, resp.StatusCode)
    }

  }
}
