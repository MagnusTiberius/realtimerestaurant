package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	ApiAddr = "35.225.142.51:8090"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Table struct {
	Code     string `json:"code"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/reservation/set/{numpeople}", ReserveGetHandler).Methods("GET")
	//r.HandleFunc("/reservation/set/{numpeople}/{datetime}", ReserveHandler).Methods("POST")
	//r.HandleFunc("/reservation/set/{numpeople}/{datetime}/{duration}", ReserveHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8088", "http://35.226.247.163:8088/"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8094", handler)
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: "", Code: 200}
	JsonResponseWrite(w, response, 200)
}

//ReserveGetHandler return a table for them to sit.
func ReserveGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	count := vars["numpeople"]
	if count == "" {
		log.Print("count undefined")
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	t := Table{Code: "1", Location: "B10", Status: "Available"}
	rv, err := json.Marshal(t)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(rv)

}
