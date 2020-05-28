package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

type contextKey string

//UserContext .
const UserContext = contextKey("user")

//ArticleHandler .
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "test category")
}

//ExampleMiddleWare .
func ExampleMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		ctx := context.WithValue(r.Context(), UserContext, "ashadi")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
