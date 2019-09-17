package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	mem := User{"Alex", 10}
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":9124", router))
}

type User struct {
	Name string
	Age  int
}
