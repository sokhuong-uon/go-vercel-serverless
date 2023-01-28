package api

import (
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog2Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		title: "2 haha",
		body:  "2 wow",
	}

	fmt.Fprintf(w, "%v", blog)
}
