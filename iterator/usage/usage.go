package main

import (
	"github.com/daniloanp/gophers-rj-meetup/iterator"
	"fmt"
)

func main() {
	var container = new(iterator.Container)
	container.Add(1)
	container.Add(2)
	container.Add("Gophers RJ")
	container.Add([]float64{.5, .6, .5612})

	var iterable iterator.Iterable = container

	for item := range iterable.Iterator() {
		switch item.(type) {
			case string:
				fmt.Printf("%q\n", item.(string))
			default:
				fmt.Println(item)
		}

	}
}

