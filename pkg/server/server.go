package server

import (
	"fmt"
	"net/http"

	"github.com/jpynicolas/hello-world/pkg/config"
)

type HttpServer struct {
	server http.Server
	mux    *http.ServeMux
}

func NewHttpServer(cnf config.Config) *HttpServer {
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", cnf.Host, cnf.Port),
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) HandlerFunc (pattern string, handler func(w http.ResponseWriter, r *http.Request)){
	s.mux.HandleFunc(pattern, handler)
}

func (s *HttpServer) StartServer() error {
	s.server.Handler = s.mux
	return s.server.ListenAndServe()
}
