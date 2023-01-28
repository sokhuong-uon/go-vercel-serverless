package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	blogs := []api.Blog{
		{"title": "Blog 1", "body": "Body"},
		{"title": "Blog 2", "body": "Body"},
		{"title": "Blog 3", "body": "Body"},
	}
	data, _ := json.Marshal(blogs)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
