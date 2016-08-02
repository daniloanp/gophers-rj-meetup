package usage

import "github.com/daniloanp/gophers-rj-meetup/observer"

type Observer1 struct {}
func (o *Observer1) Update(o Observable, arguments ...observer.Argument) {
	
}

type Observer2 struct {}
func (o *Observer2) Update(o Observable, arguments ...observer.Argument) {

}

type Observer3 struct {}
func (o *Observer3) Update(o Observable, arguments ...observer.Argument) {

}

type Observable struct {
	notifyCount int
	observer.BaseObservable
}

func (o *Observable) Notify() {
	o.NotifyObservers(o.notifyCount)
}


func main() {

}