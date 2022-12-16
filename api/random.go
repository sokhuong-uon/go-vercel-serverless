package api

import (
	"fmt"
	"math/rand"
	"net/http"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {

  number := rand.ExpFloat64()

  fmt.Fprintf(w, "%v", number)
}
