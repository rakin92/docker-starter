package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/julienschmidt/httprouter"
)

var update = flag.Bool("update", false, "update .golden files")

// TestGetDatas tests all GET REST APIs
func TestGetDatas(t *testing.T) {

	// Create Server
	s, err := NewServer()
	if err != nil {
		t.Fatalf("Server Error: %+v\n", err)
	}

	// Create cases
	cases := []struct {
		testName     string
		serverMethod httprouter.Handle
		httpMethod   string
		urlPath      string
	}{
		{"getData1", s.getData1, "GET", "/getData1"},
		{"getData2", s.getData2, "GET", "/getData2"},
	}

	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {

			// Record request to server
			req := httptest.NewRequest(tt.httpMethod, tt.urlPath, nil)
			w := httptest.NewRecorder()
			tt.serverMethod(w, req, nil)

			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)

			// Test common things

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Status Code: %#v != http.StatusOK: %#v\n", resp.StatusCode, http.StatusOK)
			}

			contentType := resp.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Content-Type: %#v != 'application/json'", contentType)
			}

			// Test body from golden file

			testDir := "testdata"

			golden := filepath.Join(testDir, tt.testName+".golden.json")
			if *update {
				err := os.MkdirAll(testDir, 0755)
				if err != nil {
					t.Fatalf("Mkdir Failure: %#v\n", err)
				}
				ioutil.WriteFile(golden, body, 0644)
			} else {
				if _, err := os.Stat(testDir); os.IsNotExist(err) {
					t.Fatalf("Run `go test -update` to create the test data")
				}
			}

			expected, err := ioutil.ReadFile(golden)
			if err != nil {
				if err != nil {
					t.Fatalf("ReadFile Failure: %#v\n", err)
				}
			}

			if !bytes.Equal(body, expected) {
				t.Logf("Actual:\n%#v", body)
				t.Logf("Expected:\n%#v", expected)
				t.Fail()
			}
		})
	}
}
