package main

import (
	"fmt"

	"github.com/santosh1608/filter"
	"github.com/santosh1608/sample-app/play"
)

func main() {
	fmt.Println("Sample app")
	filter.Filter()
	Run()
	play.Play()
}

func Run() {
	fmt.Println("running")
	filter.Filter()
}
