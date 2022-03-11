package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	output := `{"status": "Connection successful"}`
	var expected, actual map[string]interface{}
	err := json.Unmarshal([]byte(output), &expected)
	if err != nil {
		t.Fatal("healthCheck: unable to encode string to JSON", err)
		return
	}
	req, err := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	HealthCheck(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	if w.Code != http.StatusOK {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("the expected json: %s is different from actual %s", expected, actual)
	}
}
