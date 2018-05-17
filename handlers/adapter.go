package handlers

import(
	"net/http"
	"igo/middleware"
)

func DefaultAdapter(handle middleware.Handler) http.Handler{
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		requester := middleware.Requester{}
        handle(requester, rw, r)
	})
}