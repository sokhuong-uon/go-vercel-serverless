package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog1Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		Title: "First Blog",
		Body:  "Wow! this is amazing",
	}
	j, _ := json.Marshal(blog)
	data := string(j)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
