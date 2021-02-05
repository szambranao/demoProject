package main

import (
	"encoding/json"
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
	constructAndSendResponse(w, "Hello world!")
}

// constructAndSendResponse makes headers for endpoint requests and
// assembles response in JSON format.
func constructAndSendResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		errCode := http.StatusInternalServerError
		errMsg := fmt.Sprintf("Failed to encode response as JSON: %s", err.Error())
		http.Error(w, errMsg, errCode)
		return
	}
}
