package handlers

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "If my armor breaks, I fuse it back together...\n")
}
