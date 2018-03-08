package main

import (
	"encoding/json"
	"io/ioutil"
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

type Restaurant struct {
	Code     string `json:"code"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

type EnrolRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	TableCount string `json:"tablelecount"`
}
type EnrolReply struct {
	ID string `json:"id"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/enrolment/request/", ReqEnrolHandler).Methods("POST")

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

//ReqEnrolHandler return a table for them to sit.
func ReqEnrolHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	var req EnrolRequest
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &req)

	t := EnrolReply{ID: "1111"}
	rv, err := json.Marshal(t)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(rv)

}
