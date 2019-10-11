package main

import (
	"fmt"
	"log"
	"net/http"
)

type server string

const text = "server example from github\n"

func (s server) Serve(w http.ResponseWriter, req *http.Request) {
	log.Printf(text)
	fmt.Fprintf(w, text)
}

// exported
var Server server
