package main

import "fmt"
import . "ex1_results"

func main() {
	fmt.Println(CountTypes("val1", 3.2, int32(3), 6, "val5", 5, make([]string, 0)))
	fmt.Println(HelloWorld())
}
