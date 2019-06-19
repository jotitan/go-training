package ex4_results

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// wait 1 second before returning time
func wait1Second()string{
	time.Sleep(time.Second)
	return time.Now().String()
}

func wait1SecondWithChanel()chan string{
	response := make(chan string,1)
	go func() {
		timeChanelLimiter2 <- struct{}{}
		time.Sleep(time.Second)
		response <-time.Now().String()
		<-timeChanelLimiter2
	}()
	return response
}

func stopServer(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	localServer.Shutdown(ctx)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

var timeChanelLimiter = make(chan struct{},3)
var timeChanelLimiter2 = make(chan struct{},3)

func getTime(w http.ResponseWriter, r *http.Request) {
	timeChanelLimiter<-struct{}{}
	t := wait1Second()
	w.Write([]byte(t))

	<-timeChanelLimiter
}

type product struct {
	id string
	quantity int
}

var productsBase = map[string]float32 {
	"1":1.5,
	"2":4,
	"3":5,
	"4":2,
	"5":2.2,
	"6":11,
	"7":7.6,
	"8":0.4,
	"9":8,
	"10":3.9,
}

func computePrice(id string, quantity int) float32{
	time.Sleep(200*time.Millisecond)

	if price,exist := productsBase[id] ; exist {
		return price * float32(quantity)
	}
	return 0
}

// Compute the total price of basket
func computeBasket(w http.ResponseWriter, r *http.Request) {
	// List of id products separates by ,
	listProducts := strings.Split(r.FormValue("products"),",")
	// List of quantity of each product separates by ,
	products := make([]product,0,len(listProducts))
	for i,q := range strings.Split(r.FormValue("quantities"),",") {
		if quantity,err := strconv.ParseInt(q,10,32) ; err == nil {
			products = append(products,product{listProducts[i],int(quantity)})
		}
	}
	resultChan := make(chan float32)

	sum := float32(0)
	// Compute in goroutine if sum was a long procesus
	waiter := sync.WaitGroup{}
	waiter.Add(len(products))
	go func(){
		for range products {
			sum+=<-resultChan
			waiter.Done()
		}
	}()

	for _, prod := range products {
		go func(p product) {
			resultChan <- computePrice(p.id, p.quantity)
		}(prod)
	}

	waiter.Wait()

	/* Compute no thread, wait end of launch to compute *
	for range products {
		sum+=<-resultChan
	}
	 */

	w.Write([]byte(fmt.Sprintf("%f",sum)))
}

func getTimeWithLimit(w http.ResponseWriter, r *http.Request) {
	response := wait1SecondWithChanel()
	select {
	case value := <-response :w.Write([]byte(value))
	case <-time.NewTimer(1500*time.Millisecond).C:
		w.WriteHeader(500)
		w.Write([]byte("Timeout"))
	}

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
	server.HandleFunc("/getTime", getTime)
	server.HandleFunc("/getTimeWithLimit", getTimeWithLimit)
	server.HandleFunc("/computeBasket", computeBasket)

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
