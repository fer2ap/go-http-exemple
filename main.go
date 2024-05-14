package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	ADDR = "ADDR"
)

func main() {

	addr, err := os.LookupEnv(ADDR)

	if !err {
		log.Fatal("Couldn't find target address")
	}

	http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { return }))

	fmt.Println("Listening to port: ", addr)

	return
}
