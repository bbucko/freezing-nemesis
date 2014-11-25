package main

import (
	"fmt"
	"github.com/bbucko/freezing-nemesis/heroku"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	host := heroku.GetEnv("HOST", "")
	port := heroku.GetEnv("PORT", "8080")
	mongoHost := heroku.GetEnv("MONGOHQ_URL", "localhost:27017")

	session, err := mgo.Dial(mongoHost)
	if err != nil {
		log.Fatal("Mongo: ", err)
	}
	defer session.Close()

	bind := fmt.Sprintf("%s:%s", host, port)
	//	log.Println("Starting server on", bind)

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("hello world")
}
