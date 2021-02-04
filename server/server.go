package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define server structure
type Server struct {
	port int
}

// Create new server function
func NewServer(port int) *Server {
	return &Server{port}
}

// Start server function
func (s *Server) Start() {
	// Request routings
	http.HandleFunc("/hello", s.helloHandler)
	// Server standup
	log.Printf("Listening on port %d....", s.port)
	portAddr := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}

// helloHandler responds to "/hello" request
func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
