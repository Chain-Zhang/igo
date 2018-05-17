package main

import(
	"igo/router"
	"igo/handlers"
)

func main(){
	router.Get("/default", handlers.DefaultAdapter(handlers.Default))
	router.Run()
}