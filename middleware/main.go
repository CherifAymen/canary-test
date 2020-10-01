package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

//Sentiment : A struct that represents our json object received from Flusk Server
type Sentiment struct {
	Sentence string  `json:"sentence"`
	Polarity float32 `json:"polarity"`
	Version  string  `json:"version"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handleRequests() {
	fmt.Println("* Running on http://localhost:8080/ (Press CTRL+C to quit)")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/testHealth", HealthHandler)
	myRouter.HandleFunc("/sentiment", sentimentHandler).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func sentimentHandler(w http.ResponseWriter, r *http.Request) {
	saLogicAPIURL := os.Getenv("URL")
	if saLogicAPIURL == "" {
		saLogicAPIURL = "http://localhost:5000"
	}
	fullURL := saLogicAPIURL + "/analyse/sentiment"

	enableCors(&w)

	// get the body of our POST request and send it to saLogic
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(reqBody) != 0 {
	
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)


		if n == 0 {
			code = http.StatusInternalServerError
			msg = "ERROR: Something, somewhere, went wrong!\n"
			w.WriteHeader(code)
			io.WriteString(w, msg)
		} else {
			
			//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
			req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			t := time.Now()
			clock := fmt.Sprintf("[%d-%02d-%02d  %02d:%02d:%02d]",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fmt.Println(req.Host+" -- "+clock+" \""+req.Method, r.URL.Path, " ", resp.StatusCode, " - ")

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}

			var sentiment Sentiment
			json.Unmarshal(body, &sentiment)
			sentiment.Version = "go"
			json.NewEncoder(w).Encode(sentiment)
		}
	}
}

// HealthHandler returns a succesful status and a message.
// For use by Consul or other processes that need to verify service health.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
}

func main() {
	handleRequests()
}
