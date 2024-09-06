package main

import (
	cp "hung1299/go-projects/caching-proxy/cmd/cachingproxy"
	"log"
)

func main() {
	c := cp.InitializeCommand()
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}