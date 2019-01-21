package ex4_results

import (
	"context"
	"fmt"
	"net/http"
)

func stopServer(w http.ResponseWriter, r *http.Request) {
	go localServer.Shutdown(context.Background())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

var localServer *http.Server

//RunServer launch an http server
func RunServer(port int) {
	server := http.NewServeMux()
	server.HandleFunc("/stop", stopServer)
	server.HandleFunc("/helloworld", helloWorld)

	localServer = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: server}
	fmt.Println("Run on port", port)
	localServer.ListenAndServe()

	fmt.Println("Server is now shutdown")
}
