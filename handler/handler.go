package handler

import (
	"io"
	"log/slog"
	"net/http"

	r "github.com/secnops/proxit/requester"
)

func Proxy(w http.ResponseWriter, req *http.Request) {
	// we will proxy everything except the port, the remote address
	// and the path due to framework stuff
	// if the requests sends a remote-address we will point the proxied request to there
	// else 127.0.0.1 it is
	pathToRequest := req.Header.Get("path")
	port := req.Header.Get("port")
	if /*pathToRequest == "" ||*/ port == "" {
		_, err := io.WriteString(w, `
		The request should be something like:
		GET /PATH HTTP/1.1
		Host: host
		Port: [port to connect to]
		Remote-address: [if you do not want to connect to a local service]
		Tls:  [if the other service is running over Tls]
		`)
		if err != nil {
			panic(err)
		}
		return
	}
	remoteAddr := req.Header.Get("remote-address")
	tls := req.Header.Get("tls")
	schema := "http://"
	if remoteAddr == "" {
		remoteAddr = "127.0.0.1"
	}
	if tls != "" {
		schema = "https://"
	}

	url := schema + remoteAddr + ":" + port + pathToRequest

	slog.Info("Request info", "Method", req.Method, "Url", url)
	i := r.Request(req.Method, url, "", map[string]string{})
	_, err := io.WriteString(w, i)
	if err != nil {
		panic(err)
	}
	return
}
