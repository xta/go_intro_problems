package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var datastore = make(map[string]string)

func main() {
	http.HandleFunc("/records", handleRecordRequest)
	http.HandleFunc("/records/count", handleCountRequest)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*  ------------------------------------
    common
    ------------------------------------
*/

func parseKeyValue(body []byte) (key string, value string) {
	input := string(body[:])
	split_input := strings.Split(input, "=")
	key, value = split_input[0], split_input[1]
	return
}

/*  ------------------------------------
    /records endpoint
    ------------------------------------
*/

func handleRecordRequest(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case "PUT":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal("Error. Unable to read request body:", err)
		}
		response := handlePutRequest(body)
		io.WriteString(w, response)

	case "GET":
		params := req.URL.Query()
		key := params["key"][0]
		response := handleGetRequest(key)
		io.WriteString(w, response)

	case "DELETE":
		params := req.URL.Query()
		key := params["key"][0]
		response := handleDeleteRequest(key)
		io.WriteString(w, response)
	}
}

func handlePutRequest(body []byte) string {
	key, value := parseKeyValue(body)
	addItem(key, value)

	var response bytes.Buffer
	response.WriteString("Success. Added ")
	response.WriteString(key)
	response.WriteString(" - ")
	response.WriteString(value)
	return response.String()
}

func handleGetRequest(key string) string {
	value, ok := lookupItem(key)
	var response bytes.Buffer

	if ok {
		response.WriteString("Success. The value for your key '")
		response.WriteString(key)
		response.WriteString("' is: ")
		response.WriteString(value)
	} else {
		response.WriteString("Error. Unable to find your key: ")
		response.WriteString(key)
	}
	return response.String()
}

func handleDeleteRequest(key string) string {
	deleteAction := deleteItem(key)
	var response bytes.Buffer
	if deleteAction {
		response.WriteString("Success. Your record was destroyed.")
	} else {
		response.WriteString("Error. Unable to delete record since your key was not found.")
	}
	return response.String()
}

/*  ------------------------------------
    /records/count endpoint
    ------------------------------------
*/

func handleCountRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, "Count: ")

		params := req.URL.Query()
		if len(params) > 0 {
			query := params["q"][0] // Warning: This does not handle keys besides "q"
			queryCount := numKeyStartsWith(query)
			io.WriteString(w, queryCount)
		} else {
			io.WriteString(w, datastoreCount())
		}
	}
}

func datastoreCount() string {
	return strconv.Itoa(len(datastore))
}

func numKeyStartsWith(query string) string {
	count := 0
	for key, _ := range datastore {
		if strings.HasPrefix(key, query) {
			count++
		}
	}
	return strconv.Itoa(count)
}

/*  ------------------------------------
    datastore
    ------------------------------------
*/

func addItem(key string, value string) {
	datastore[key] = value
}

func lookupItem(key string) (value string, ok bool) {
	value, ok = datastore[key]
	return
}

func deleteItem(key string) (existence bool) {
	existence = keyExists(key)
	if existence {
		delete(datastore, key)
	}
	return
}

func keyExists(key string) (ok bool) {
	_, ok = datastore[key]
	return
}
