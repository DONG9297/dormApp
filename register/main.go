package main

import (
	"net/http"

	"register/controller"
)

func main() {
	http.HandleFunc("/register", controller.Register)
	http.ListenAndServe(":10711", nil)
}
