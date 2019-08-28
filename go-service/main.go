package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NormalizedColumn is the json structure
type NormalizedColumn struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// NewServer returns a server to serve from
func NewServer() (*Server, error) {
	data1 := []NormalizedColumn{
		{1, "one"},
		{2, "two"},
	}
	data2 := []NormalizedColumn{
		{3, "three"},
		{4, "four"},
	}
	return &Server{data1, data2}, nil
}

// Server that has data to use
type Server struct {
	data1 []NormalizedColumn
	data2 []NormalizedColumn
}

func (s *Server) getData1(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ") // NOTE: pretty-printing might not be good on an API
	e.Encode(s.data1)

}

func (s *Server) getData2(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ") // NOTE: pretty-printing might not be good on an API
	e.Encode(s.data2)

}

func main() {
	s, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GET("/getData1", s.getData1)
	router.GET("/getData2", s.getData2)

	log.Fatal(http.ListenAndServe(":3000", router))
}
