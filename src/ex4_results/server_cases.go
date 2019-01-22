package ex4_results

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func stopServer(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	localServer.Shutdown(ctx)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func add5ToValue(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("value")
	if valueAsInt, err := strconv.ParseInt(value, 10, 32); err == nil {
		w.Write([]byte(fmt.Sprintf("%d", valueAsInt+5)))
		return
	}
	http.Error(w, "Bad parameter for value", 403)
}

var localServer *http.Server

//RunServer launch an http server and link functions on rest endpoints
func RunServer(port int) {
	server := http.NewServeMux()
	server.HandleFunc("/stop", stopServer)
	server.HandleFunc("/helloworld", helloWorld)
	server.HandleFunc("/add5", add5ToValue)

	localServer = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: server}
	fmt.Println("Run on port", port)
	localServer.ListenAndServe()

	fmt.Println("Server is now shutdowned")
}

//CallURL call an url, return result as string and status code
func CallURL(url string) (string, int) {
	if resp, err := http.Get(url); err == nil {
		data, _ := ioutil.ReadAll(resp.Body)
		return string(data), resp.StatusCode
	} else {
		return "", 404
	}
}
