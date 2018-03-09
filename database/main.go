package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
	r.HandleFunc("/database/create/reservation", createReservationTable).Methods("GET")

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

//createReservationTable func
func createReservationTable(w http.ResponseWriter, r *http.Request) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-central1-a")},
	)

	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Code"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Code"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Reservation"),
	}

	_, err = svc.CreateTable(input)
	//awsconfig := &aws.Config{
	//	Region: aws.String("us-central1-a"),
	//}
	//svc := dynamodb.New(nil, awsconfig)
	//_, err := tables.CreateReservationTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)
	response := ResponseMessage{Message: "Reservation table created", Code: 200}
	JsonResponseWrite(w, response, 200)

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
