package observer

type Argument interface {}
type Observer interface {
	Update(o Observable, arguments ...Argument)
}
