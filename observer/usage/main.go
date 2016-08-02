package main

import (
	"github.com/daniloanp/gophers-rj-meetup/observer"
	"time"
	"fmt"
	"math/rand"
)

type Observer1 struct {}
func (_ *Observer1) Update(o observer.Observable, arguments ...observer.Argument) {
	for { // looping funcion, a problematic one
		time.Sleep(12 * time.Second)
		fmt.Println(arguments)
		fmt.Println("==> 1",arguments)
	}
}

type Observer2 struct {}
func (_ *Observer2) Update(o observer.Observable, arguments ...observer.Argument) {
	time.Sleep(5 * time.Second)
	fmt.Println("==> 2",arguments)
}

type Observer3 struct {}
func (_ *Observer3) Update(o observer.Observable, arguments ...observer.Argument) {
	fmt.Println("===> 3",arguments)
}

type Observable struct {
	notifyCount int
	observer.BaseObservable
}

func (o *Observable) Notify() {
	o.NotifyObservers(o.notifyCount)
	o.notifyCount += 1
}


func main() {
	observable := &Observable{}
	observers := []observer.Observer {
		new(Observer1),
		new(Observer2),
		new(Observer3),
	}

	for _,o := range observers {
		observable.AddObserver(o)
	}

	fmt.Println("Beginning cycle!")

	for {
		sleepTime := time.Duration(1 + rand.Intn(40)%2) * time.Second
		fmt.Println("Sleep time:", sleepTime)
		time.Sleep(sleepTime)
		observable.Notify()
	}

}