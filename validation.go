package validation

import (
	"github.com/wlMalk/mirango/framework"
)

type Validator interface {
	Validate(framework.Context, framework.ParamValue) error
	Description() string
}

type ValidatorFunc func(framework.Context, framework.ParamValue) error

func (f ValidatorFunc) Validate(c framework.Context, v framework.ParamValue) error {
	return f(c, v)
}

func (f ValidatorFunc) Description() string {
	return ""
}

type FuncValidator struct {
	function    func(framework.Context, framework.ParamValue) error
	description string
}

func NewFuncValidator(f func(framework.Context, framework.ParamValue) error, d string) Validator {
	return &FuncValidator{
		function:    f,
		description: d,
	}
}

func (f FuncValidator) Validate(c framework.Context, v framework.ParamValue) error {
	return f.function(c, v)
}

func (f FuncValidator) Description() string {
	return f.description
}
