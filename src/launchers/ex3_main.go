package main

import (
	"fmt"
	"github.com/jotitan/go-training/ex3"
	"time"
)

func main() {
	begin := time.Now()
	fmt.Println(ex3.ComputeSum(1000000))
	fmt.Println("Total time", time.Now().Sub(begin))
}
