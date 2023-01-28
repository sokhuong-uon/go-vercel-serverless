package api

import (
	"fmt"
	"net/http"
)

type Blog struct {
	Title string `Json:"title"`
	Body  string `Json:"body"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey! Let's chat")
}
