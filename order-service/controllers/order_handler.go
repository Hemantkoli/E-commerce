package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Hemantkoli/e-commerence/order-service/models"
	"github.com/Hemantkoli/e-commerence/order-service/utils"
)

func RemoveOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w, http.StatusBadRequest, "Invalid data :(")
		return
	}
	order := &models.Order{}
	err = order.RemoveOrder(uint(id))
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "failed"
		response["message"] = "Error while deleting the order :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "Order deleted successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}

func RemoveAllOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w, http.StatusBadRequest, "Invalid data :(")
		return
	}
	order := &models.Order{}
	err = order.RemoveAllOrders(uint(userId))
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "failed"
		response["message"] = "Error while deleting the orders :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "Orders deleted successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w,
			http.StatusBadRequest,
			map[string]string{"error": "Error while creating order :("})
		return
	}
	response := make(map[string]interface{})
	if order.AddOrder() != nil {
		response["result"] = "failed"
		response["message"] = "Error while creating order :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "Order placed successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}

func FetchAllOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w, http.StatusBadRequest, "Invalid data :(")
		return
	}
	order := &models.Order{}
	userOrders, err := order.FetchAllOrders(uint(userId))
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "failed"
		response["message"] = "Error while fetching orders for the user :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		utils.ResponseJson(w, http.StatusOK, userOrders)
	}
}