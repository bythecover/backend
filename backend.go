package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/poll_events/", func(w http.ResponseWriter, r *http.Request) {
		newString := strings.TrimPrefix(r.URL.Path, "/poll_events/")
		val, _ := strconv.Atoi(newString)
		fmt.Fprintf(w, "%d", val)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
