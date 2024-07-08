package main

import (
	"fmt"
	"log"
	"net/http"
)

// API
func main() {

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	http.HandleFunc("/health", handleHealth)

	log.Println("Serving on localhost:4242")

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("ENPOINT CALLED")

}

func handleHealth(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("HEALTH CHECK")
	var response []byte = []byte("Server is up and running")
	writer.Write(response)

}
