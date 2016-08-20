package validation

import (
	"regexp"

	"github.com/mirango/framework"
)

var (
	RegexpNumber = regexp.MustCompile("^[0-9]+$")
	RegexpFloat  = regexp.MustCompile("^[+-]{0,1}([0-9]*[.])?[0-9]+$")
	RegexpAlpha  = regexp.MustCompile("^[a-zA-Z]+$")
)

func Eq(a interface{}) Validator {
	return NewFuncValidator(func(c framework.Context, v framework.ParamValue) error {
		if a != v.Value() {
			return ErrEq.Err([]interface{}{v.Name(), a}...)
		}
		return nil
	}, MsgEq.Msg(a))
}

func Number() Validator {
	return Regexp(RegexpNumber)
}

func Alpha() Validator {
	return Regexp(RegexpAlpha)
}

func Regexp(p *regexp.Regexp) Validator {
	return NewFuncValidator(func(c framework.Context, v framework.ParamValue) error {
		if !p.MatchString(v.RawString()) {
			return ErrRegexp.Err([]interface{}{v.Name(), p.String()}...)
		}
		return nil
	}, MsgRegexp.Msg(p.String()))
}

func StringRegexp(pstr string) Validator {
	p, err := regexp.Compile(pstr)
	if err != nil {
		return nil
	}
	return Regexp(p)
}
