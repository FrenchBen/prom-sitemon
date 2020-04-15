package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.String())
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "PUT" && r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var page AlertManagerData
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	bodyText, err := ioutil.ReadAll(r.Body)
	if err == nil {
		log.Printf("%s\n", bodyText)
	} else {
		log.Println("Error reading body")
		return
	}
	if string(bodyText) == "" {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		log.Printf("Body is empty JSON: %s\n", bodyText)
		return
	}
	err = json.Unmarshal(bodyText, &page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("JSON Decode Error: %v\n", err.Error())
		return
	}
	if len(page.Alerts) == 0 {
		http.Error(w, "No alerts to display", http.StatusBadRequest)
		log.Println("No alerts to display")
		return
	}
	log.Println("Descripion: " + page.Alerts[0].Annotations.Description)
	log.Println("Status: " + page.Status)
	w.Write([]byte("OK"))
}

// e.g. http.HandleFunc("/_health", healthCheckHandler)
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/_health" {
		http.NotFound(w, r)
		return
	}
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	w.Write([]byte(`{"alive": true}`))
}

func main() {
	/// Set up logging to output file
	logfile := os.Getenv("LOG_ALERT_PATH")
	if logfile != "" {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("Error opening log file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/_health", healthCheckHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}
