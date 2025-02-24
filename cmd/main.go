package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/secnops/proxit/handler"
)

func main() {

	exposePort := flag.String("exposePort", "8000", "port to expose")
	var b strings.Builder
	b.WriteString(":")
	b.WriteString(*exposePort)
	localAddr := b.String()
	flag.Parse()

	http.HandleFunc("/", handler.Proxy)
	log.Fatal(http.ListenAndServe(localAddr, nil))
}
