package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Flags.
	addr := flag.String("addr", ":5555", "server address:port")
	abspath, _ := filepath.Abs(".")
	dir := flag.String("dir", abspath, "directory to serve from")
	flag.Parse()

	// Create the server object.
	srv := http.FileServer(http.Dir(*dir))

	// Tell the user what's going on.
	log.Printf("serving from %q.", *dir)
	log.Printf("listening on %q.", *addr)

	// Serve.
	log.Fatal(http.ListenAndServe(*addr, srv))
}
