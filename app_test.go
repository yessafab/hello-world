package main

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	var tests = []struct {
		Description        string
		URL                string
		Method             string
		Body               string
		ExpectedStatusCode int
		ExpectedBody       string
	}{
		{
			Description:        "Expected use case",
			URL:                "/",
			Method:             "GET",
			ExpectedStatusCode: 200,
			ExpectedBody:       "hello world",
		},
	}
	for _, test := range tests {
		bodyReq := bytes.NewBufferString(test.Body)
		req := httptest.NewRequest(test.Method, test.URL, bodyReq)
		resp := httptest.NewRecorder()

		rootHandle(resp, req)
		// check status code correct
		if resp.Code != test.ExpectedStatusCode {
			t.Errorf("%s - Status Code Fail. Expected: '%d' Got: '%d'", test.Description, test.ExpectedStatusCode, resp.Code)
		}
		bodyRespBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("%s - Response Body failed to read all with error: %s", test.Description, err.Error())
		}
		bodyResp := string(bodyRespBytes)
		if bodyResp != test.ExpectedBody {
			t.Errorf("%s - Response Body incorrect. Expected: '%s' Got: '%s'", test.Description, test.ExpectedBody, bodyResp)
		}
	}
	return
}
