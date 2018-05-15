package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"math/rand"

	. "app.nazul/config"
	. "app.nazul/repo"
	"app.nazul/service"
	"github.com/gorilla/mux"
)

var processRequests = 0

func ServerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := map[string]interface{}{
		"appName":    CONFIG.AppName,
		"appVersion": CONFIG.AppVersion,
		"host":       CONFIG.Host,
		"port":       CONFIG.Port,
		"startTime":  CONFIG.Start.Format("2006-01-02 15:04:05"),
		"processed":  processRequests,
	}
	json.NewEncoder(w).Encode(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processRequests++
		time := time.Now().Format("20060102150405")
		randId := rand.Intn(10000)
		seq := fmt.Sprintf("%02d-%s-%04d", CONFIG.AppInstanceId, time, randId)
		w.Header().Set("SEQ", seq)
		u, p, ok := r.BasicAuth()
		next.ServeHTTP(w, r)
		if CONFIG.AccessLog {
			accMap := map[string]interface{}{
				"appname":       CONFIG.AppName,
				"instanceid":    CONFIG.AppInstanceId,
				"seq":           seq,
				"rhost":         r.Host,
				"rheader-type":  r.Header["Content-type"],
				"rheader-agent": r.Header["User-Agent"],
				"rmethod":       r.Method,
				"url":           r.URL.Path,
				"rlength":       r.ContentLength,
				"rauth":         fmt.Sprintf("%s/%s-%v", u, p, ok),
			}
			logJson, _ := json.Marshal(accMap)
			log.Printf("%s", string(logJson))
		}
	})
}

func main() {
	defer CONN.Close()
	log.Printf("Using Config:\n\n%v\n\n", CONFIG)
	r := mux.NewRouter()
	r.HandleFunc("/", ServerStatus).Methods("GET", "POST")
	http.HandleFunc("/", ServerStatus)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	/***RSA START***/
	r.HandleFunc("/api/rsa/{bits}", service.GenerateKeys).Methods("GET")
	r.HandleFunc("/api/rsa", service.RandomKeyPair).Methods("GET")
	/***RSA END***/
	/***USERS START***/
	r.HandleFunc("/api/users", service.CreateUser).Methods("POST")
	r.HandleFunc("/api/users", service.UsersPage).Methods("GET")
	r.HandleFunc("/api/users/{id}", service.UserUpdate).Methods("PUT")
	r.HandleFunc("/api/users/{id}", service.UserByIdOrName).Methods("GET")
	r.HandleFunc("/api/users/{id}", service.UserDelete).Methods("DELETE")
	// r.HandleFunc("/api/users/list", service.UsersList).Methods("GET")
	/***USERS END***/
	/***BOOKS START***/
	r.HandleFunc("/api/books", service.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/list", service.GetBooks).Methods("GET", "POST")
	r.HandleFunc("/api/books/{id}", service.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", service.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", service.DeleteBook).Methods("DELETE")
	/***BOOKS END***/

	r.Use(loggingMiddleware)
	log.Printf("Server running at port:%s:%d", CONFIG.Host, CONFIG.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", CONFIG.Host, CONFIG.Port), r))
}
