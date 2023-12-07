package main

import (
	"fmt"
	"log"
	"net/http"
)

var isOnline = true

func main() {
	http.HandleFunc("/api/greeting", greeting)
	http.HandleFunc("/api/request", requestUrl)
	http.HandleFunc("/api/stop", stop)
	http.HandleFunc("/api/health", health)
	http.Handle("/", http.FileServer(assetFS()))
	fmt.Println("Web server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestUrl(w http.ResponseWriter, r *http.Request) {
	//https://vault-vault.apps.cluster-blmgx.blmgx.sandbox322.opentlc.com/
	if isOnline {
		requestUrl := "www.google.com"
		if u := r.FormValue("url"); u != "" {
			requestUrl = u
		}

		_, err := http.Get(requestUrl)
		if err != nil {
			log.Printf("Error during Get is: %s", err) // throw 509
			fmt.Fprintf(w, "Call to [%s] failed", err)
		} else {
			log.Printf("Call is done")
			fmt.Fprintf(w, "Call to [%s] performed!", requestUrl)
		}

		return
	}
	w.WriteHeader(503)
	w.Write([]byte("Not Online"))

}

func greeting(w http.ResponseWriter, r *http.Request) {
	if isOnline {
		message := "World"
		if m := r.FormValue("name"); m != "" {
			message = m
		}
		fmt.Fprintf(w, "Hello %s!", message)
		return
	}
	w.WriteHeader(503)
	w.Write([]byte("Not Online"))
}

func stop(w http.ResponseWriter, r *http.Request) {
	isOnline = false
	w.Write([]byte("Stopping HTTP Server"))
}

func health(w http.ResponseWriter, r *http.Request) {
	if isOnline {
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(500)
}
