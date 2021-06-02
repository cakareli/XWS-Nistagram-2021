package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/cakareli/XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
)

func main() {
	var aa model.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello from AUTHENTICATION SERVICE!")
	})


	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
