package api

import (
	"encoding/json"
	"net/http"

	"github.com/hozzer/apitemplate/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/user", s.handleGetUserByID)
	// http.HandleFunc()
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO
	user, err := s.store.Get("1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
