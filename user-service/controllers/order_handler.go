package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Hemantkoli/e-commerence/user-service/utils"
)

const (
	FETCH_ALL_ORDERS_URL = "orders/user/"
	ADD_ORDER_URL       = "order"
	REMOVE_ALL_ORDERS_URL = "orders/user/"
	REMOVE_ORDER_URL   = "order/"
)

func RemoveAllOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVICE_URL") + REMOVE_ALL_ORDERS_URL + vars["userId"])
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body := utils.ResponseJsonBody(resp)
	utils.ResponseJson(w, http.StatusOK, body)
}

func RemoveOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVICE_URL") + REMOVE_ORDER_URL + vars["orderId"])
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body := utils.ResponseJsonBody(resp)
	utils.ResponseJson(w, http.StatusOK, body)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(os.Getenv("ORDER_SERVICE_URL") + ADD_ORDER_URL)
	request, err := http.NewRequest(http.MethodPost, url, r.Body)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body := utils.ResponseJsonBody(resp)
	utils.ResponseJson(w, http.StatusOK, body)
}

func FetchAllOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVICE_URL") + FETCH_ALL_ORDERS_URL + vars["userId"])
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body := utils.ResponseJsonBody(resp)
	utils.ResponseJson(w, http.StatusOK, body)
}