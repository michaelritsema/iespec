package main

import (
	"fmt"
	"net/http"
)

func zflowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good job")
	fmt.Println(r)

}

func main() {
	http.HandleFunc("/api/concentrator", zflowHandler)
	http.ListenAndServe(":8443", nil)
}
