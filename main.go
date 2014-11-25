package main

import (
	"fmt"
	"github.com/bbucko/freezing-nemesis/heroku"
	"github.com/yvasiyarov/gorelic"
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
	newRelic := heroku.GetEnv("NEW_RELIC_LICENSE_KEY", "")

	agent := gorelic.NewAgent()
	agent.Verbose = true
	agent.NewrelicLicense = newRelic
	agent.Run()

	log.Println("Connecting to mongo: ", mongoHost)
	session, mongoErr := mgo.Dial(mongoHost)

	if mongoErr != nil {
		log.Fatal("Mongo: ", mongoErr)
	}
	defer session.Close()

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("Starting server on", bind)

	http.HandleFunc("/", handler)
	httpErr := http.ListenAndServe(bind, nil)
	if httpErr != nil {
		log.Fatal("ListenAndServe: ", httpErr)
	}
	fmt.Println("hello world")
}
