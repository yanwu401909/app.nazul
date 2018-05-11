package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"math/rand"

	. "app.nazul/config"
	"app.nazul/service"
	"github.com/gorilla/mux"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/yaml.v2"
)

var cfg Config = Config{}
var processRequests = 0
var esClient elastic.Client

var esMapping = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"_default_" : {
			"properties":{
				"appname":{
					"type":"keyword"
				},
				"instanceid":{
					"type":"integer"
				},
				"seq":{
					"type":"keyword"
				},
				"rhost":{
					"type":"ip"
				},
				"rheader-type":{
					"type":"keyword"
				},
				"rheader-agent":{
					"type":"keyword"
				},
				"rmethod":{
					"type":"keyword"
				},
				"rlength":{
					"type":"integer"
				},
				"rauth":{
					"type":"keyword"
				}
			}
		}
	}
}
`

func ServerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data := map[string]interface{}{
		"appName":    cfg.AppName,
		"appVersion": cfg.AppVersion,
		"host":       cfg.Host,
		"port":       cfg.Port,
		"startTime":  cfg.Start.Format("2006-01-02 15:04:05"),
		"processed":  processRequests,
	}
	json.NewEncoder(w).Encode(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processRequests++
		time := time.Now().Format("20060102150405")
		randId := rand.Intn(10000)
		seq := fmt.Sprintf("%02d-%s-%04d", cfg.AppInstanceId, time, randId)
		w.Header().Set("SEQ", seq)
		u, p, ok := r.BasicAuth()
		next.ServeHTTP(w, r)

		logJson := map[string]interface{}{
			"appname":       cfg.AppName,
			"instanceid":    cfg.AppInstanceId,
			"seq":           seq,
			"rhost":         r.Host,
			"rheader-type":  r.Header["Content-type"],
			"rheader-agent": r.Header["User-Agent"],
			"rmethod":       r.Method,
			"rlength":       r.ContentLength,
			"rauth":         fmt.Sprintf("%s/%s-%v", u, p, ok),
		}
		index, err := esClient.Index().
			Index(cfg.Elasticsearch.EsIndex).
			Type(cfg.Elasticsearch.EsType).
			Id(seq).
			BodyJson(logJson).
			Do(context.Background())
		if err != nil {
			log.Printf("Waring:es log error! %v", err)
		}
		log.Printf("%v", index)

	})
}

func initESClient(cfg Config) {
	url := fmt.Sprintf("http://%s:%d", cfg.Elasticsearch.Host, cfg.Elasticsearch.Port)
	esClient, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf("Waring:es client init error! %v", err)
	}
	esversion, err := esClient.ElasticsearchVersion(url)
	if err != nil {
		log.Printf("Waring:es version fetch error! %v", err)
	}
	fmt.Printf("Elasticsearch  version [%s]\n\n", esversion)

	exists, err := esClient.IndexExists(cfg.Elasticsearch.EsIndex).Do(context.Background())
	if err != nil {
		log.Printf("Waring: es IndexExists error! %v", err)
	}
	if !exists {
		createIndex, err := esClient.CreateIndex(cfg.Elasticsearch.EsIndex).Body(esMapping).Do(context.Background())
		if err != nil {
			log.Printf("Waring:es create index [%s] error! %v", cfg.Elasticsearch.EsIndex, err)
		}
		log.Printf("Result:%v", createIndex)
		if !createIndex.Acknowledged {

		}
	}
}

func main() {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Start = time.Now()
	log.Printf("Using Config:\n\n%v\n\n", cfg)
	initESClient(cfg)
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", ServerStatus).Methods("GET", "POST")
	/***USERS START***/

	/***USERS END***/
	/***BOOKS START***/
	r.HandleFunc("/api/books", service.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", service.GetBook).Methods("GET")
	r.HandleFunc("/api/books", service.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", service.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", service.DeleteBook).Methods("DELETE")
	/***BOOKS END***/
	r.Use(loggingMiddleware)
	log.Printf("Server running at port:%s:%d", cfg.Host, cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), r))
}
