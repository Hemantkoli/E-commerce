package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/Hemantkoli/e-commerence/user-service/controllers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.FetchAllUsers).Methods("GET")
	router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	router.HandleFunc("/user", controllers.RemoveUser).Methods("DELETE")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/order", controllers.AddOrder).Methods("POST")
	router.HandleFunc("/user/{userId}/orders", controllers.FetchAllOrders).Methods("GET")
	router.HandleFunc("/user/{userId}/orders", controllers.RemoveAllOrders).Methods("DELETE")
	router.HandleFunc("/user/{userId}/order/{orderId}", controllers.RemoveOrder).Methods("DELETE")

	port := os.Getenv("USER_SERVICE")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("User service running on port:", port)
}