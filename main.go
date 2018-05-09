package main
import (
	"fmt"
	"net/http"
	"log"
)

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
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":8000", mux))
}