package mydevservice

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dev host=%v\n", r.Host)
}
