package traces

type trace struct {
	debug bool
}

func (t *trace) SetDebug(isDebug bool) {
	t.debug = isDebug
}
////////////////////
type FileTrace struct {
	trace
}

func (t *FileTrace) Debug(msg string) {
	if t.debug {
		// print msg
	}
}

func (t *FileTrace) Error(msg string) {
	//print msg
}
////////////////////


type ServiceTrace struct {
	trace
}

func (t *ServiceTrace) Debug(msg string) {
	if t.debug {
		// print msg
	}
}

func (t *ServiceTrace) Error(msg string) {
	//print msg
}


const somecond bool = false

func Trace



