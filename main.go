package main

import (
	"fmt"

	"github.com/santosh1608/filter"
	"github.com/santosh1608/sample-app/dance"
	"github.com/santosh1608/sample-app/play"
)

func main() {
	fmt.Println("Sample app")
	filter.Filter()
	Run()
	play.Play()
	dance.Dance()
}

func Run() {
	fmt.Println("running")
	filter.Filter()
}
