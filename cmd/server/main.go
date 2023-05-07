package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpynicolas/hello-world/pkg/config"
	"github.com/jpynicolas/hello-world/pkg/server"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	server := server.NewHttpServer(cnf.Server)
	server.HandlerFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	fmt.Printf("starting server on %s:%d", cnf.Server.Host, cnf.Server.Port)
	if err := server.StartServer(); err != nil{
		log.Println("error starting server")
	}
}
