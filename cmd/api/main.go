package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	srv := &http.Server{
		Addr:    ":4000",
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	fmt.Println(err)

}
