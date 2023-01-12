package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The total should not equal to Zero")
	assert.Equal(t, 4, total, "Expecting 4")
}

func TestSubtract(t *testing.T) {
	total := Subtract(1, 3)
	assert.NotNil(t, total, "The total should not equal to Zero")
	assert.Equal(t, -2, total, "Expecting -2")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok response")
	assert.Equal(t, "Hello World", response.Body.String(), "incorrect body")

}
