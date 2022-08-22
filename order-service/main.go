package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/Hemantkoli/e-commerence/order-service/controllers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/order", controllers.AddOrder).Methods("POST")
	router.HandleFunc("/orders/user/{userId}", controllers.FetchAllOrders).Methods("GET")
	router.HandleFunc("/order/{id}", controllers.RemoveOrder).Methods("DELETE")
	router.HandleFunc("/orders/user/{userId}", controllers.RemoveAllOrders).Methods("DELETE")

	port := os.Getenv("ORDER_SERVICE")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Order service running on port:", port)
}