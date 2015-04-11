package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func echoIp(w http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	io.WriteString(w, host)
}

func main() {
	httpAddr := flag.String("http", ":3999", "HTTP service address (e.g., ':3999')")
	flag.Parse()

	http.HandleFunc("/ip", echoIp)
	err := http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
