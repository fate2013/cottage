package server

import (
	"fmt"
	"net/http"

	serverlib "github.com/nicholaskh/golib/server"
)

type CottageServer struct {
	*serverlib.HttpServer
}

func NewCottageServer() *CottageServer {
	this := new(CottageServer)
	this.HttpServer = serverlib.NewHttpServer("cottage")
	this.initRouter()
	return this
}

func (this *CottageServer) initRouter() {
	this.HttpServer.RegisterHandler("/", this.test)
}

func (this *CottageServer) test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

func (this *CottageServer) Launch(listenAddr string) {
	this.HttpServer.Launch(listenAddr)
}
