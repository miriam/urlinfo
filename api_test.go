package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestCases = [][2]string{
	{"google.com", "true"},
	{"google.com/", "false"},

	{"google.com/foobar", "true"},
	{"google.com/foobar/", "false"},

	{"google.com/frob", "false"},
	{"google.com/frob/", "true"},

	{"yahoo.com?123=4", "true"},
	{"yahoo.com?123=4&f=x&.43owe?^%$0f", "true"},

	{"yahoo.com", "false"},
	{"yahoo.com/a/b/c/d", "false"},
	{"yahoo.com/a/b/c/d/3", "true"},
	{"go.hotbot.eu/2.3/4.5/?x=y", "true"},

	{"example.com:443", "true"},
	{"example.com:444", "false"},
	{"example.com:444/", "true"},
	{"example.com:443/1/2/3?a=b", "true"},
}

func TestBlocklistBlocksListMembers(t *testing.T) {
	router := setupRouter()

	for i := 0; i < len(TestCases); i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/urlinfo/1/%s", TestCases[i][0]), nil)
		router.ServeHTTP(w, req)

		expectedBody := fmt.Sprintf("{\"blocklisted\":%s}", TestCases[i][1])

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	}
}
