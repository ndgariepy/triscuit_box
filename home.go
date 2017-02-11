package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TrisKitBox, the best place on the web to track your wagers")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	appengine.Main()
}
