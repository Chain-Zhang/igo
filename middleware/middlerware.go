package middleware

import(
	"fmt"
	"net/http"
)

type Requester struct {
	Mobile       string   // 用户手机号
	Email        string
	Name         string   // 用户名称
	ID           string   // 用户唯一标示
}

type Handler func(requester Requester, w http.ResponseWriter, r *http.Request)

func Middleware1(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("arrive at Middleware1")
        next.ServeHTTP(w, r)
    })
}

func Middleware2(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("arrive at Middleware2")
        next.ServeHTTP(w, r)
    })
}