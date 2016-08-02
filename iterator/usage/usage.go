package iterator

import (
	"github.com/daniloanp/gophers-rj-meetup/iterator"
	"fmt"
)

func main() {
	var container iterator.Iterable
	container = new(iterator.Container)

	for item := range container.Iterator() {
		fmt.Println(item)
	}
}