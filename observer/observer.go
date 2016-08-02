package observer

type Argument interface {}
type Observer interface {
	Update(Observable,  ...Argument)
}
