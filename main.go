package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello from inside the docker")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello darker !!!</h1>")
	})

	log.Fatal(http.ListenAndServe(":1082", nil))
}
