package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	b := []byte(`{"username":"foo","password":"bar"}`)
	payload := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodGet, "/signup", payload)
	if err != nil {
		t.Fail()
	}
	res := httptest.NewRecorder()

	Signup(res, req)
	got := res.Body.String()
	if len(got) == 0 {
		t.Errorf("expected token")
	}
}
