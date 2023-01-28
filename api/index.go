package api

import (
	"fmt"
	"net/http"
)

type Blog struct {
	title string
	body  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey! Let's chat")
}
