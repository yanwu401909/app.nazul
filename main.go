package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app.nazul/constant"
	"app.nazul/models"
	"app.nazul/service"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := new(models.EchoResponse)
	resp.ResultCode = constant.OK
	resp.ResultMessage = constant.CODE_MAPPING[constant.OK]
	resp.Data = r.URL.Path
	if data, err := json.Marshal(resp); err == nil {
		fmt.Fprintf(w, "%s", data)
	} else {
		fmt.Fprint(w, "{}")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := new(models.EchoResponse)
	loginname := r.FormValue("loginname")
	password := r.FormValue("password")
	repass := r.FormValue("repass")
	if loginname == "" || password == "" || repass == "" {
		resp.ResultCode = constant.PARAMS_ERROR
		resp.ResultMessage = constant.CODE_MAPPING[constant.PARAMS_ERROR]
	} else if password != repass {
		resp.ResultCode = constant.REPASS_ERROR
		resp.ResultMessage = constant.CODE_MAPPING[constant.REPASS_ERROR]
	} else {
		resp.ResultCode = constant.OK
		resp.ResultMessage = constant.CODE_MAPPING[constant.OK]
	}
	if data, err := json.Marshal(resp); err == nil {
		fmt.Fprintf(w, "%s", data)
	} else {
		fmt.Fprint(w, "{}")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	path := r.URL.Path
	if path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Site!<h1>")
	} else if path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <br/><a href=\"mailto:yanwu401909@gmail.com\">yanwu401909@gmail.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :( </h1><br/><a href='/'>BACK HOME</a>")
	}
}

func main() {
	log.Println("Server run at port:8000")
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	r.HandleFunc("/", Index)
	r.HandleFunc("/api/regist", Regist).Methods("GET", "POST")
	r.HandleFunc("/api/login", Login).Methods("GET", "POST")
	r.HandleFunc("/api/logout", Logout).Methods("GET", "POST")
	r.HandleFunc("/api/user/{uid}", User).Methods("GET", "POST")
	/***BOOKS START***/
	r.HandleFunc("/api/books", service.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", service.GetBook).Methods("GET")
	r.HandleFunc("/api/books", service.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", service.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", service.DeleteBook).Methods("DELETE")
	/***BOOKS END***/

	log.Fatal(http.ListenAndServe(":8000", r))
}
