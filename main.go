package main
import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"app.nazul/models"
)

const(
	
) 

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	resp := new(models.EchoResponse)
	resp.ResultCode = OK
	resp.ResultMessage="success"
	resp.Data = r.URL.Path
	if data, err := json.Marshal(resp); err == nil{
		fmt.Fprintf(w,"%s", data)
	}else{
		fmt.Fprint(w, "{}")
	}
}

func Regist(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	resp := new(models.EchoResponse)
	loginname := r.FormValue("loginname")
	password := r.FormValue("password")
	repass := r.FormValue("repass")
	if loginname == "" || password == "" || repass == ""{
		resp.ResultCode = PARAMS_ERROR
	}else if password != repass{
		resp.ResultCode = REPASS_ERROR
	}else{
		resp.ResultCode = OK
	}
	if data, err := json.Marshal(resp); err == nil{
		fmt.Fprintf(w,"%s", data)
	}else{
		fmt.Fprint(w, "{}")
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func Logout(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func User(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w, "{\"resultCode\":0,\"resultMessage\":\"Success\"}")
}

func handlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
	path := r.URL.Path
	if path == "/"{
		fmt.Fprint(w, "<h1>Welcome to my Site!<h1>")
	}else if path == "/contact"{
		fmt.Fprint(w, "To get in touch, please send an email to <br/><a href=\"mailto:yanwu401909@gmail.com\">yanwu401909@gmail.com</a>.")
	}else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"<h1>We could not find the page you were looking for :( </h1><br/><a href='/'>BACK HOME</a>")
	}
}

func main(){
	log.Println("Server run at port:8000")
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	r.HandleFunc("/", Index)
	r.HandleFunc("/regist", Regist).Methods("GET", "POST")
	r.HandleFunc("/login", Login).Methods("GET", "POST")
	r.HandleFunc("/logout", Logout).Methods("GET", "POST")
	r.HandleFunc("/user/{uid}", User).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}