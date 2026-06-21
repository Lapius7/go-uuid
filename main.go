package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	fmt.Fprintln(w, id.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7100", nil)
}
