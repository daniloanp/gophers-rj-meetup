package factory_method


type Trace interface {
	SetDebug(bool)
	Debug(msg string)
	Error(msg string)
}

//TraceFactory interface we should embedd this
type TraceFactory interface {
	Trace() Trace
}

//
//var factory  = TraceFactory(nil)
//
//func CreateTrace() Trace {
//	if factory == nil {
//		return nil
//	}
//	return factory.Trace()
//}
//
//func SetFactory(f TraceFactory) {
//	factory = f
//}



