package main

import (
	"fmt"
	"net/http"

	Router "mygo/router"
)

func main() {
	fmt.Println("Start")
	r := Router.NewRouter()
	http.ListenAndServe(":3001", r)
}
