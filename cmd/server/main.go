package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.Int("p", 8080, "port to listen")
	confpath := flag.String("c", "", "cluster config path")

	if err := Serve(*port, GetConfig(*confpath)); err != nil {
		log.Fatal(err)
	}
}
