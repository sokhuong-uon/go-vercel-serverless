package api

import (
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	blogs := []api.Blog{
		{Title: "Blog 1", Body: "Body"},
		{Title: "Blog 2", Body: "Body"},
		{Title: "Blog 3", Body: "Body"},
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", blogs)
}
