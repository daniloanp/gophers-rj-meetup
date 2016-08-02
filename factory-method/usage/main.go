package main

import "github.com/daniloanp/gophers-rj-meetup/factory-method"

type Application interface {
	Run()
	Tracer() factory_method.Trace
}

func (b Application) Run() {
	b.Tracer().Debug("Running")
}

type BaseApplication struct {
	name string
	code int64
}



type RealApplication interface





func main() {

}
