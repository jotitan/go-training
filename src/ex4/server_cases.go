package ex4

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//stopServer stop the running server.
func stopServer(w http.ResponseWriter, r *http.Request) {
	go localServer.Shutdown(context.Background())
}

var localServer *http.Server

// wait 1 second before returning time
func wait1Second()string{
	return ""
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

	w.Write([]byte("0"))

}

//RunServer launch an http server
func RunServer(port int) {}

//CallURL call an url, return result as string and status code
func CallURL(url string) (string, int) {
	return "", 0
}
