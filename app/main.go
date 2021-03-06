package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func word2Number(s string) string {
	s = strings.ToLower(s)
	if s == "" {
		return "0"
	}
	var res string
	for i := 0; i < len(s); i++ {
		c := byte(s[i])
		if c >= 97 && c <= 123 {
			res = res + strconv.Itoa(int(s[i])-96)
		} else if c >= 48 && c <= 57 {
			res = res + strconv.Itoa(int(s[i])-48)
		}
	}
	return res
}

func getWord2Number(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := make(map[string]string)
	response["result"] = word2Number(vars["param"])
	json.NewEncoder(w).Encode(response)
	fmt.Println("Endpoint Hit -> getWord2Number")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["status"] = "ok"
	json.NewEncoder(w).Encode(response)
	fmt.Println("Endpoint Hit -> getStatus")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ready", getStatus).Methods("GET")
	router.HandleFunc("/{param}", getWord2Number).Methods("GET")

	log.Println("Listening on port 4200...")
	log.Fatal(http.ListenAndServe(":4200", router))
}
