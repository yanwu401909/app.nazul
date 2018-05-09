package main
import (
	"fmt"
	"net/http"
	"log"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to my site!!!</h1>")
}
func main(){
	log.Println("Server run at port:8000")
	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":8000", nil))
}