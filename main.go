package main

import (
	"net/http"
	authcontroller "github.com/LutfiEkaprima/Goproject/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
}