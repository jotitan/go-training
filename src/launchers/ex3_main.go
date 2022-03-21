package main

import (
	"fmt"
	"github.com/jotitan/go-training/ex3_results"
	"time"
)

func main() {
	begin := time.Now()
	fmt.Println(ex3_results.ComputeSum(1000000))
	fmt.Println("Total time", time.Now().Sub(begin))
}
