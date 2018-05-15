package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "app.nazul/errors"
	. "app.nazul/models"
	"app.nazul/repo"
	"github.com/gorilla/mux"
)

func UserByIdOrName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]
	data := UserResponse{ApiResponse: ApiResponse{OK, "success"}}
	u, err := repo.FindUserByIdOrLoginName(id)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	} else if u.Id == "" {
		data.ResultCode = NOTEXIST_ERROR
		data.ResultMessage = CODE_MAPPING[NOTEXIST_ERROR]
	} else {
		data.Data = u
	}
	json.NewEncoder(w).Encode(data)
}

func UsersPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := UsersPageResponse{ApiResponse: ApiResponse{OK, "success"}}
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 || pageSize > 50 {
		pageSize = 10
	}
	page, err := repo.UsersPage(pageNo, pageSize)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	} else {
		data.Data = page
	}
	json.NewEncoder(w).Encode(data)
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := UsersListResponse{ApiResponse: ApiResponse{OK, "success"}}
	users, err := repo.UsersList()
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	} else {
		data.Data = users
	}
	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := ApiResponse{OK, "success"}
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Status = 1
	user.CreateTime = time.Now()
	err := repo.SaveUser(&user)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]
	data := ApiResponse{OK, "success"}
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := repo.UpdateUser(id, &user)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]
	data := ApiResponse{OK, "success"}
	err := repo.DeleteUser(id)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}
