package main

import (

	"fmt"
	"log"
	"net/http"

)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])

	})


	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLO GEEKS!")

	})
	fmt.Println("Listening on 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
