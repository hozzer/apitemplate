package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hozzer/apitemplate/api"
	"github.com/hozzer/apitemplate/storage"
)

func main() {
	listenAddr := flag.String("listen", ":8080", "address to listen on")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("starting server on", *listenAddr)
	log.Fatal(server.Run())
}
