package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/secnops/proxit/handler"
)

func main() {

	exposePort := flag.String("exposePort", "8000", "port to expose")
	flag.Parse()
	var b strings.Builder
	b.WriteString(":")
	b.WriteString(*exposePort)
	localAddr := b.String()

	server := &http.Server{
		Addr:              localAddr,
		ReadHeaderTimeout: 5 * time.Second,
	}
	http.HandleFunc("/", handler.Proxy)
	log.Fatal(server.ListenAndServe())
}
