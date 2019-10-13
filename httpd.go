// Trivial HTTP-server which allows you to specify a host/port/directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//
	// Parse command-line flags.
	//
	path := flag.String("path", ".", "The directory to use as the HTTP root directory")
	addr := flag.String("host", "127.0.0.1", "The host to bind upon (use 0.0.0.0 for remote access)")
	port := flag.Int("port", 3000, "The port to listen upon")

	flag.Parse()

	//
	// Create a static-file server, based upon the
	// path we're treating as our root-directory.
	//
	fs := http.FileServer(http.Dir(*path))
	http.Handle("/", fs)

	//
	// Build up the listen addres.
	//
	listen := fmt.Sprintf("%s:%d", *addr, *port)

	//
	// Log our start, and begin serving.
	//
	log.Printf("Serving upon http://%s/\n", listen)
	http.ListenAndServe(listen, logRequest(http.DefaultServeMux))
}

// logRequest dumps the request to the console.
//
// Of course we don't know the return-code, but this is good enough
// for most of my use-cases.
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
