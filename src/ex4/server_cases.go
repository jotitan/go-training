package ex4

import (
	"context"
	"net/http"
)

func stopServer(w http.ResponseWriter, r *http.Request) {
	go localServer.Shutdown(context.Background())
}

var localServer *http.Server

//RunServer launch an http server
func RunServer(port int) {}
