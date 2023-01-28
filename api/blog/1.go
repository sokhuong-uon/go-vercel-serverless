package api

import (
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog1Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		title: "Haha",
		body:  "wow",
	}

	fmt.Fprintf(w, "%v", blog)
}
