package factory_method

import (
	"io"
)

type Trace interface {
	SetDebug(bool)
	Debug(msg string)
	Error(msg string)
}

type trace struct {
	io.Writer
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
