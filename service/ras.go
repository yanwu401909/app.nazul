package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"app.nazul/repo"

	. "app.nazul/errors"
	. "app.nazul/models"
)

func GenerateKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := ApiResponse{OK, "success"}
	bits := mux.Vars(r)["bits"]
	if bits != "1024" && bits != "2048" {
		data.ResultCode = PARAMS_ERROR
		data.ResultMessage = CODE_MAPPING[PARAMS_ERROR]
	} else {
		bits, _ := strconv.Atoi(bits)
		err := repo.GenerateRsaKey(bits)
		if err.Code != OK {
			data.ResultCode = SERVER_ERROR
			data.ResultMessage = CODE_MAPPING[SERVER_ERROR]
		}
	}
	json.NewEncoder(w).Encode(data)
}

func RandomKeyPair(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := RsaPairResponse{ApiResponse: ApiResponse{OK, "success"}}
	pair, err := repo.RandomRsaPair()
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	} else {
		data.Data = pair
	}
	json.NewEncoder(w).Encode(data)
}
