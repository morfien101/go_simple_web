package main

import(
	"os"
	"net/http"
	"io"
	"fmt"
	"time"
)

const VERSION = "0.0.1"

func responder(w http.ResponseWriter, r *http.Request) {
	hostname,_ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("<html><head><title>%s</title></head><body><h1>You hit: %s<br />Version:%s</h1></body></html>",hostname,hostname,VERSION))
	fmt.Println(time.Now().UTC(),"/")
}

func main() {
	http.HandleFunc("/", responder)
	http.ListenAndServe(":8000",nil)
}
