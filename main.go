package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	// base file dir
	base string
	// server addr
	addr string
	// custom response headers
	headers string
)

func init() {
	flag.StringVar(&base, "base", "./data", "tell me where to find files")
	flag.StringVar(&addr, "addr", "127.0.0.1:18888", "tell me which host and port to listen on")
	flag.StringVar(&headers, "headers", "", "tell me your custom response headers, format like 'Field1:Value1//Field2:Value2//Field3:Value3'")
}

func main() {
	flag.Parse()
	if base == "" {
		log.Fatalf("you have to give me a directory")
	}
	if addr == "" {
		log.Fatalf("you have to give me a addr")
	}
	hdrs := strings.Split(strings.TrimSpace(headers), "//")
	http.Handle("/", respHeaderswrapper(http.FileServer(http.Dir(base)), hdrs))
	http.HandleFunc("/301test", func(r http.ResponseWriter, req *http.Request) {
		r.Header().Set("Location: ", "https://www.example.com/")
		r.WriteHeader(http.StatusMovedPermanently)
	})
	log.Printf("serving %s on addr: http://%s\n", base, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func respHeaderswrapper(hdlr http.Handler, headers []string) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(headers); i++ {
			if h := headers[i]; h != "" {
				s := strings.Split(h, ":")
				if len(s) != 2 {
					continue
				}
				w.Header().Add(s[0], s[1])
			}
		}
		hdlr.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
