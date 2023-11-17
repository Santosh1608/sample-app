package main

import (
	"fmt"

	"github.com/santosh1608/filter"
)

func main() {
	fmt.Println("Sample app")
	filter.Filter()
	Run()
}

func Run() {
	fmt.Println("running")
	filter.Filter()
}
