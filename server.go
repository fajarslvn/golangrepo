package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/fajarslvn/go_rest_api/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000" 

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "up and running...")
	})

	router.HandleFunc("/posts", route.GetPosts).Methods("GET")
	router.HandleFunc("/posts", route.AddPost).Methods("POST")
	log.Println("listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
} 