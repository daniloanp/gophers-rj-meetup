package traces

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

