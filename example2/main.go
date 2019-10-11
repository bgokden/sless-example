package main

import (
	"fmt"
	"log"
	"net/http"
)

type server string

const text = "server example from github 2\n"

func (s server) Name() string {
	return "example2"
}

func (s server) Serve(w http.ResponseWriter, req *http.Request) {
	log.Printf(text)
	fmt.Fprintf(w, text)
}

// exported
var Server server
