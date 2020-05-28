package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestArticleHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	res := httptest.NewRecorder()

	//use this method if you need mux variable in your code
	// r := mux.NewRouter()
	// r.HandleFunc("/test", ArticleHandler)
	// r.ServeHTTP(res, req)

	//use this method to test simple handler if you don't need any mux variable
	ArticleHandler(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "test category", string(body))
}

func TestExampleMiddleware(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(UserContext)
		user, ok := val.(string)
		assert.Equal(t, true, ok)
		assert.Equal(t, "ashadi", user)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, user)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	res := httptest.NewRecorder()

	//use simple method
	// bHandler := ExampleMiddleWare(testHandler)
	// bHandler.ServeHTTP(res, req)

	//use mux router
	m := mux.NewRouter()
	m.Use(ExampleMiddleWare)
	m.HandleFunc("/test", testHandler)
	m.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "ashadi", string(body))

}
