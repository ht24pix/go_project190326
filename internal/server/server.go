// file: internal/server/server.go
package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, welcome to the Warehouse Management System!")
	})
	http.ListenAndServe(":8080", nil)
}
