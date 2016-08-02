package traces


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