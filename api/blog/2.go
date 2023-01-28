package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog2Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		"title": "Second Blog",
		"body":  "ok",
	}
	data, _ := json.Marshal(blog)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
