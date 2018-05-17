package router

import(
	"net/http"
)

type Route struct{
	Method string
	Path string
	Action http.Handler
}

var routes []Route

func Get(path string, action http.Handler){
    routes = append(routes, Route{Method:"GET", Path:path, Action:action})
}