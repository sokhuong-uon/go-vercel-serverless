package handler

import (
  "fmt"
  "net/http"
  "math/rand"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {

  number := rand.ExpFloat64()

  fmt.Fprintf(w, "<h1>Random number: %v</h1>",number)
}
