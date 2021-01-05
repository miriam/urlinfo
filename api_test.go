package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

var TestCases = [5][2]string {
	{"google.com", "true"},
	{"google.com/foobar", "true"},
	{"yahoo.com?123=4", "true"},
	{"yahoo.com", "false"},
	{"yahoo.com/a/b/c/d", "false"},
}

func TestBlocklist(t *testing.T) {
	router := setupRouter()

	for i := 0; i < len(TestCases); i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/urlinfo/1/%s", TestCases[i][0]), nil)
		router.ServeHTTP(w, req)

		expectedBody := fmt.Sprintf("{\"blocklisted\":%s,\"originalUrl\":\"%s\"}", TestCases[i][1], TestCases[i][0])
	
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	}
}

