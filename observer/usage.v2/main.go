package main

import (
	"github.com/daniloanp/gophers-rj-meetup/observer/observer.v2"
	"time"
	"fmt"
	"math/rand"
)

type Observer1 struct {
	Channel observer.Observer
}
func (obs *Observer1) Listen() {
	for {
		message := <-obs.Channel
		go func() {
			for { // looping funcion, a problematic one
				time.Sleep(12 * time.Second)
				fmt.Println("==> 1",message.Arguments)
			}
		}()
	}
}

type Observer2 struct {
	Channel observer.Observer
}
func (obs *Observer2) Listen() {
	for {
		message := <-obs.Channel
		go func() {
			for { // looping funcion, a problematic one
				time.Sleep(5 * time.Second)
				fmt.Println("==> 2",message.Arguments)
			}
		}()
	}
}

type Observer3 struct {
	Channel observer.Observer
}
func (obs *Observer3) Listen() {
	for {
		message := <-obs.Channel
		fmt.Println("===> 3",message.Arguments)
	}
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
	o1 := &Observer1{make(observer.Observer)}
	o2 := &Observer2{make(observer.Observer)}
	o3 := &Observer3{make(observer.Observer)}


	observable.AddObserver(o1.Channel)
	observable.AddObserver(o2.Channel)
	observable.AddObserver(o3.Channel)

	go func() { o1.Listen() }()
	go func() { o2.Listen() }()
	go func() { o3.Listen() }()

	fmt.Println("Beginning cycle!")

	for {
		sleepTime := time.Duration(1 + rand.Intn(40)%2) * time.Second
		fmt.Println("Sleep time:", sleepTime)
		time.Sleep(sleepTime)
		observable.Notify()
	}

}