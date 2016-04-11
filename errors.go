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

type Err string

func (err Err) Err(vals ...interface{}) error {
	return ValidationErr(fmt.Sprintf(string(err), vals...))
}

type Msg string

func (m Msg) Msg(vals ...interface{}) string {
	return fmt.Sprintf(string(m), vals...)
}

type ValidationErr string

func (err ValidationErr) Error() string {
	return string(err)
}

const (
	ErrEq     = Err("%v must equal %v") // prefix "Validation Error: " will be added
	ErrRegexp = Err("%v must match %v")
)

const (
	MsgEq     = Msg("must equal %v") // prefix "Validation Error: " will be added
	MsgRegexp = Msg("must match %v")
)
