package server

import "net/http"

type Server struct {
	Addr    string
	Handler http.Handler
}

func NewServer(p string, h http.Handler) *Server {
	return &Server{
		Addr:    p,
		Handler: h,
	}

}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	return srv.ListenAndServe()
}
