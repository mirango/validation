package validation

import (
	"fmt"
)

type err struct {
	Msg    string
	Values map[string]interface{}
}

func (e *err) Error() string {
	return e.Msg
}

type Error map[string][]error

func (e Error) Error() string {
	return ""
}

func (e Error) Append(name string, err ...error) {
	e[name] = append(e[name], err...)
}

func (e Error) Set(name string, err []error) {
	e[name] = err
}

func (e Error) Union(err Error) {
	for k := range err {
		e[k] = err[k]
	}
}

func (e Error) UnionAppend(err Error) {
	for k := range err {
		e[k] = append(e[k], err[k]...)
	}
}
