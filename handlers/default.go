package handlers

import(
	"fmt"
	"net/http"
	"igo/middleware"
)

func Default (requester middleware.Requester, rw http.ResponseWriter, r *http.Request){
	fmt.Fprintf(rw, "this is default")
}