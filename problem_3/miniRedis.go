package main

import (
	"io"
	"log"
	"net/http"
)

var datastore map[string]string

func main() {
	http.HandleFunc("/records/counts", handleCountRequest)
	http.HandleFunc("/records", handleRecordRequest)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleRecordRequest(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "records")
}

func handleCountRequest(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "counts")
}
