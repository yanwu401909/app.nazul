package service

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"app.nazul/mock"
	"app.nazul/models"
	"app.nazul/constant"
	"math/rand"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	data := new(models.BooksResponse)
	data.ResultCode = constant.OK
	data.ResultMessage = constant.CODE_MAPPING[constant.OK]
	data.Data = mock.BOOKS_DATA
	json.NewEncoder(w).Encode(data)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	id := mux.Vars(r)["id"]
	data := new(models.BookResponse)
	for _,item := range mock.BOOKS_DATA{
		if item.Id == id{
			data.ResultCode = constant.OK
			data.ResultMessage = constant.CODE_MAPPING[constant.OK]
			data.Data = item
			json.NewEncoder(w).Encode(data)
			return
		}
	}
	data.ResultCode = constant.NOTEXIST_ERROR
	data.ResultMessage = constant.CODE_MAPPING[constant.NOTEXIST_ERROR]
	json.NewEncoder(w).Encode(data)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	data := new(models.ApiResponse)
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	if book.Title == "" || book.Isbn == "" {
		data.ResultCode = constant.PARAMS_ERROR
		data.ResultMessage = constant.CODE_MAPPING[constant.PARAMS_ERROR]
	}else {
		book.Id = strconv.Itoa(rand.Intn(100000) + 10)
		mock.BOOKS_DATA = append(mock.BOOKS_DATA, book)
		data.ResultCode = constant.OK
		data.ResultMessage = constant.CODE_MAPPING[constant.OK]
	}
	json.NewEncoder(w).Encode(data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	id := mux.Vars(r)["id"]
	data := new(models.ApiResponse)
	data.ResultCode = constant.NOTEXIST_ERROR
	data.ResultMessage = constant.CODE_MAPPING[constant.NOTEXIST_ERROR]
	for index,item := range mock.BOOKS_DATA{
		if item.Id == id{
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			if book.Title == "" || book.Isbn == "" {
				data.ResultCode = constant.PARAMS_ERROR
				data.ResultMessage = constant.CODE_MAPPING[constant.PARAMS_ERROR]
			}else {
				mock.BOOKS_DATA = append(mock.BOOKS_DATA[:index], mock.BOOKS_DATA[index+1:]...)
				book.Id = id
				mock.BOOKS_DATA = append(mock.BOOKS_DATA, book)
				data.ResultCode = constant.OK
				data.ResultMessage = constant.CODE_MAPPING[constant.OK]
			}
			break
		}
	}
	json.NewEncoder(w).Encode(data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	id := mux.Vars(r)["id"]
	data := new(models.ApiResponse)
	for index,item := range mock.BOOKS_DATA{
		if item.Id == id{
			mock.BOOKS_DATA = append(mock.BOOKS_DATA[:index], mock.BOOKS_DATA[index+1:]...)
			break
		}
	}
	data.ResultCode = constant.OK
	data.ResultMessage = constant.CODE_MAPPING[constant.OK]
	json.NewEncoder(w).Encode(data)
}