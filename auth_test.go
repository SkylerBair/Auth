package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
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

	a := []byte(`{"username":"foo","password":"bar"}`)
	SignInpayload := bytes.NewReader(a)

	SignInreq, err := http.NewRequest(http.MethodGet, "/signup", SignInpayload)
	if err != nil {
		t.Fail()
	}
	SignInres := httptest.NewRecorder()

	Signin(SignInres, SignInreq)
	SignIngot := SignInres.Body.String()
	t.Logf("%v", SignIngot)
	if len(SignIngot) == 0 {
		t.Errorf("expected Token")
	}

}
