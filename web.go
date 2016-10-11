package main

import(
	"os"
	"net/http"
	"io"
	"fmt"
)

func responder(w http.ResponseWriter, r *http.Request) {
	hostname,_ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("<html><head><title>%s</title></head><body><h1>You hit: %s</h1></body></html>",hostname,hostname))
}

func main() {
	http.HandleFunc("/", responder)
	http.ListenAndServe(":8000",nil)	
}
