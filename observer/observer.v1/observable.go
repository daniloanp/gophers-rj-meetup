package observer


type Observable interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers(arguments ... Argument)
}

type BaseObservable struct {
	observers []Observer
}

//Missing some lock or cocnurrency protection
func (b *BaseObservable) AddObserver(o Observer) {
	if b.observers == nil {
		b.observers = make([]Observer, 0)
	}
	for _, item:=  range  b.observers {
		if item == o {
			return
		}
	}
	b.observers = append(b.observers, o)
}

//Missing some lock or cocnurrency protection
func (b* BaseObservable) RemoveObserver(o Observer) {
	for inx, item:=  range  b.observers {
		if item == o {
			s := b.observers
			s = append(s[:inx], s[inx+1:]...)
			b.observers = s
			return
		}
	}
}

func (b *BaseObservable) NotifyObservers(arguments ... Argument) {
	f := func(o Observer) { //do not block if observer block
		o.Update(b, arguments...)
	}
	for _, observer :=  range  b.observers {
		go f(observer)
	}
}