package api

import (
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog1Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		"title": "First Blog",
		"body":  "Wow! this is amazing",
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", blog)
}
