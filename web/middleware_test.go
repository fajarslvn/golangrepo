package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After execute Handler")

}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo executed")
		fmt.Fprint(w, "Hello Foo")
	})

	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: LogMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
