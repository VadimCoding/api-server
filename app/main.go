package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
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

func writeInDb(db *sql.DB, call map[string]string) {
	sqlStatement := `
INSERT INTO call_history (time, param, result)
VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, call["time"], call["param"], call["result"])
	if err != nil {
		panic(err)
	}
}

func getWord2Number(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := make(map[string]string)
	response["result"] = word2Number(vars["param"])
	json.NewEncoder(w).Encode(response)
	fmt.Println("Endpoint Hit -> getWord2Number")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	var call = map[string]string{"time": time.Now().String(), "param": vars["params"], "result": response["result"]}
	writeInDb(db, call)

	// close database
	defer db.Close()

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
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
