package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// API
func main() {

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	http.HandleFunc("/health", handleHealth)

	log.Println("Serving on localhost:4242")

	var err = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("ENPOINT CALLED")

	if request.Method != "POST" {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		// data that we are receiving from the client
		ProductId string `json:"product_id"`
		FirsName  string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address1"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(req)

}

func handleHealth(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("HEALTH CHECK")
	var response = []byte("Server is up and running")
	_, err := writer.Write(response) // _ means ignore the return value
	if err != nil {
		fmt.Println(err)
	}

}
