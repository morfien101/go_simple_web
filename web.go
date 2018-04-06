package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	listenPort = flag.Int("p", 8000, "Port to listen on.")
)

const VERSION = "0.0.1"

func responder(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("<html><head><title>%s</title></head><body><h1>You hit: %s<br />Version:%s</h1></body></html>", hostname, hostname, VERSION))
	fmt.Println(time.Now().UTC(), "/")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", responder)
	srvAddr := fmt.Sprintf(":%d", *listenPort)
	err := http.ListenAndServe(srvAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
