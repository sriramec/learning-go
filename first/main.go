package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello World")

	fmt.Println("Hi im there")

	fmt.Println("num of cpus is ", runtime.NumCPU())
}
