package handler
import (
  "fmt"
  "net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>so-uo: Hey!</h1>")
}
