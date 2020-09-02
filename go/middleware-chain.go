package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func chain(middlewares []middleware, handler http.HandlerFunc) http.HandlerFunc {
	fn := handler
	for _, m := range middlewares {
		fn = m(fn)
	}

	return fn
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

func recoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	middlewares := []middleware{recoverMiddleware, loggingMiddleware} // loggingMiddleware(recoverMiddleware(HandlerFunc))

	http.HandleFunc("/", chain(middlewares, myHandlerFunc))
	http.ListenAndServe("127.0.0.1:8080", nil)
}
