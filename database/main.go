package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MagnusTiberius/realtimerestaurant/database/tables"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const CREATING = "CREATING..."

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/database/create/reservation", CreateReservationTable).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8088", "http://35.226.247.163:8088/"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8097", handler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: "", Code: 200}
	JsonResponseWrite(w, response, 200)
}

func CreateReservationTable(w http.ResponseWriter, r *http.Request) {
	awsconfig := &aws.Config{
		Region: aws.String("us-central1-a"),
	}
	svc := dynamodb.New(nil, awsconfig)
	_, err := tables.CreateReservationTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)
}

func JsonResponseWrite(w http.ResponseWriter, message interface{}, statusCode int) {

	body, err := json.Marshal(message)

	if statusCode == 200 && err == nil {
		msg, _ := json.Marshal(message)
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	} else {
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, string(body), statusCode)
		}
	}
}
