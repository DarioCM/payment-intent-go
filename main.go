package main

import (
	"fmt"
	"log"
	"net/http"
)

// API
func main() {

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("ENPOINT CALLED")

}
