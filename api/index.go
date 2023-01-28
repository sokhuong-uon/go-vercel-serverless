package api

import (
	"fmt"
	"net/http"
)

type Blog map[string]string

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's chat")
}
