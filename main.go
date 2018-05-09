package main
import (
	"fmt"
	"net/http"
	"log"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
	w.Header().Set("Auth","lvye")
	fmt.Fprint(w, "<h1>Welcome to my super awesome site!!!</h1>")
}

func handlerSayHi(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	if name == ""{
		name = "lvye"
	}
	fmt.Fprint(w, "你好," + name + "!!!")
}

func main(){
	log.Println("Server run at port:8000")
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/sayHi", handlerSayHi)
	log.Fatal(http.ListenAndServe(":8000", nil))
}