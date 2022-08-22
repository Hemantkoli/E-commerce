package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hemantkoli/e-commerence/user-service/models"
	"github.com/Hemantkoli/e-commerence/user-service/utils"
)

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
	}
	err = user.RemoveUser(user.ID)
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "failed"
		response["message"] = "Error while deleting user :("
		utils.ResponseJson(w, http.StatusUnprocessableEntity, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "User deleted successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	users, err := user.FetchAllUsers()
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "failed"
		response["message"] = "Error while fetching users :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		utils.ResponseJson(w, http.StatusOK, users)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w,
			http.StatusInternalServerError,
			map[string]string{"error": "Error while updating user :("})
		return
	}
	response := make(map[string]interface{})
	if user.UpdateUser() != nil {
		response["result"] = "failed"
		response["message"] = "Error while updating user :("
		utils.ResponseJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "User updated successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		utils.ResponseJson(w,
			http.StatusInternalServerError,
			map[string]string{"error": "Error while creating user :("})
	}
	response := make(map[string]interface{})
	if user.AddUser() != nil {
		response["result"] = "failed"
		response["message"] = "Error while creating user :("
		utils.ResponseJson(w, http.StatusUnprocessableEntity, response)
	} else {
		response["result"] = "succeed"
		response["message"] = "User created successfully :)"
		utils.ResponseJson(w, http.StatusOK, response)
	}
}