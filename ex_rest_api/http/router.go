package router

import "net/http"

type Router interface {
	GET(uri string, f func(res http.ResponseWriter, req *http.Request))
	POST(uri string, f func(res http.ResponseWriter, req *http.Request))
	SERVE(port string)
}