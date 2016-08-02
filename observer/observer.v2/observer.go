package observer

type Argument interface {}

type Message struct {
	Origin Observable
	Arguments []Argument
}

type Observer chan *Message
