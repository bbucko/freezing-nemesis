package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getEnv(key string, defVal string) (env string) {
	env = os.Getenv(key)

	if env == "" {
		env = defVal
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	host := getEnv("HOST", "")
	port := getEnv("PORT", "8080")

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("Starting server on", bind)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
