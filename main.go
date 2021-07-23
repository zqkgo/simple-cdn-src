package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	// base file dir
	base string
	// server addr
	addr string
)

func init() {
	flag.StringVar(&base, "base", "./data", "tell me where to find files")
	flag.StringVar(&addr, "addr", "127.0.0.1:18888", "tell me which host and port to listen on")
}

func main() {
	flag.Parse()
	if base == "" {
		log.Fatalf("you have to give me a directory")
	}
	if addr == "" {
		log.Fatalf("you have to give me a addr")
	}
	http.Handle("/", http.FileServer(http.Dir(base)))
	log.Printf("serving %s on addr: http://%s\n", base, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
