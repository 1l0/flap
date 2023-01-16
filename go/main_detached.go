//go:build detached

package main

import (
	"github.com/1l0/flap/go/server"
	"log"
)

func main() {
	log.Println("Listen on http://localhost:9090")
	server.Serve("9090", "~/Downloads/")
}
