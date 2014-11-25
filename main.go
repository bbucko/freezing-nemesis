package main

import (
	"fmt"
	"github.com/bbucko/freezing-nemesis/heroku"
	"github.com/codegangsta/negroni"
	negronigorelic "github.com/jingweno/negroni-gorelic"
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
	newRelicLicenseKey := heroku.GetEnv("NEW_RELIC_LICENSE_KEY", "")

	log.Println("[GoRelic]", "Starting GoRelic agent with key: ", newRelicLicenseKey)
	agent := gorelic.NewAgent()
	agent.NewrelicLicense = newRelicLicenseKey
	agent.Run()

	log.Println("[Mongo]", "Connecting to mongo: ", mongoHost)
	session, mongoErr := mgo.Dial(mongoHost)

	if mongoErr != nil {
		log.Fatal("[Mongo]", mongoErr)
	}
	defer session.Close()

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("[FreezingFrenzy]", "Starting server on", bind)

	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success!\n")
	})

	n := negroni.New()

	n.Use(negronigorelic.New(newRelicLicenseKey, "freezing-nemesis", false))
	n.UseHandler(r)

	n.Run(bind)
	log.Println("[FreezingFrenzy]", "Ready to be start serving request")
}
