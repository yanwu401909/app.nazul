package service

import (
	"encoding/json"
	"net/http"

	. "app.nazul/errors"
	. "app.nazul/models"
	"app.nazul/repo"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := BooksListResponse{ApiResponse: ApiResponse{OK, "success"}}
	data.Data, _ = repo.GetBooks()
	json.NewEncoder(w).Encode(data)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]
	data := BookResponse{ApiResponse: ApiResponse{OK, "success"}}
	book, err := repo.GetBook(id)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	data.Data = book
	json.NewEncoder(w).Encode(data)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := ApiResponse{OK, "success"}
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := repo.SaveBook(book)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := ApiResponse{OK, "success"}
	id := mux.Vars(r)["id"]
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	err := repo.UpdateBook(id, book)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := ApiResponse{OK, "success"}
	id := mux.Vars(r)["id"]
	err := repo.DeleteBook(id)
	if err.Code != OK {
		data.ResultCode = err.Code
		data.ResultMessage = err.Message
	}
	json.NewEncoder(w).Encode(data)
}
