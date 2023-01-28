package api

import (
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog2Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		Title: "Second Blog",
		Body:  "A good blog post",
	}

	fmt.Fprintf(w, "%v", blog)
}
