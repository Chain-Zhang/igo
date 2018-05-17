package router

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"igo/middleware"
)

var r *mux.Router



func Run(){
	initRouter()
	fmt.Println("server start")
	fmt.Println("port: 3000")
	http.ListenAndServe(":3000", r)
}

func initRouter(){
	r = mux.NewRouter()
	r.Use(middleware.Middleware1)
	r.Use(middleware.Middleware2)
	for _, route := range routes{
		r.Handle(route.Path, route.Action).Methods(route.Method)
	}
}

