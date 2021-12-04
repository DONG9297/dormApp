package main

import (
	"login/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/is_logged", controller.IsLogged)
	http.HandleFunc("/check_password", controller.CheckPassword)
	http.ListenAndServe(":10712", nil)
}
