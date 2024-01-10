package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Authentification struct {
	username string
	passwort string
}

func main() {

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		var auth Authentification
		err = json.Unmarshal(body, &auth)
		if err != nil {
			panic(err)
		}

		log.Println(auth)

	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Listening on Port 3000")
	http.ListenAndServe(":3000", nil)
}
