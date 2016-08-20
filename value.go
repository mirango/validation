package validation

import (
	"strconv"

	"github.com/mirango/framework"
)

type Value struct {
	name             string
	value            interface{}
	strValue         string
	as               framework.ValueType
	strMultipleValue []string
	multiple         bool
	from             string
}

func NewValue(name string, value string, from string, as framework.ValueType) *Value {
	return &Value{
		name:     name,
		strValue: value,
		from:     from,
		as:       as,
	}
}

func NewMultipleValue(name string, value []string, from string, as framework.ValueType) *Value {
	return &Value{
		name:             name,
		strMultipleValue: value,
		multiple:         true,
		from:             from,
		as:               as,
	}
}

func (v *Value) Name() string {
	return v.name
}

func (v *Value) As() framework.ValueType {
	return v.as
}

func (v *Value) Value() interface{} {
	if v.value == nil {
		switch v.as {
		case framework.TYPE_STRING:
			return v.String()
		case framework.TYPE_INT:
			return v.Int64()
		case framework.TYPE_FLOAT:
			return v.Float64()
		case framework.TYPE_UINT:
			return v.Uint64()
		// case framework.TYPE_COMPLEX:
		// 	return v.Complex128()
		case framework.TYPE_BOOL:
			return v.Bool()
		default:
			return v.strValue
		}
	}
	return v.value
}

func (v *Value) RawString() string {
	return v.strValue
}

func (v *Value) String() string {
	if v.value != nil {
		rv := v.value.(string)
		return rv
	}
	if v.as == framework.TYPE_STRING {
		v.value = v.strValue
		return v.strValue
	}
	return ""
}

func (v *Value) BoolE() (bool, error) {
	if v.value != nil {
		rv, ok := v.value.(bool)
		if ok {
			return rv, nil
		}
	}
	if v.as == framework.TYPE_BOOL {
		rv, err := strconv.ParseBool(v.strValue)
		if err == nil {
			v.value = rv
		}
		return rv, err
	}
	return false, nil
}

func (v *Value) Bool() bool {
	rv, _ := v.BoolE()
	return rv
}

func (v *Value) Int64E() (int64, error) {
	if v.value != nil {
		rv, ok := v.value.(int64)
		if ok {
			return rv, nil
		}
	}
	if v.as == framework.TYPE_INT {
		rv, err := strconv.ParseInt(v.strValue, 10, 64)
		if err == nil {
			v.value = rv
		}
		return rv, err
	}
	return 0, nil
}

func (v *Value) Int64() int64 {
	rv, _ := v.Int64E()
	return rv
}

func (v *Value) Int32E() (int32, error) {
	rv, err := v.Int64E()
	return int32(rv), err
}

func (v *Value) Int32() int32 {
	return int32(v.Int64())
}

func (v *Value) Int16E() (int16, error) {
	rv, err := v.Int64E()
	return int16(rv), err
}

func (v *Value) Int16() int16 {
	return int16(v.Int64())
}

func (v *Value) Int8E() (int8, error) {
	rv, err := v.Int64E()
	return int8(rv), err
}

func (v *Value) Int8() int8 {
	return int8(v.Int64())
}

func (v *Value) IntE() (int, error) {
	rv, err := v.Int64E()
	return int(rv), err
}

func (v *Value) Int() int {
	return int(v.Int64())
}

func (v *Value) Float64E() (float64, error) {
	if v.value != nil {
		rv, ok := v.value.(float64)
		if ok {
			return rv, nil
		}
	}
	if v.as == framework.TYPE_FLOAT {
		rv, err := strconv.ParseFloat(v.strValue, 64)
		if err == nil {
			v.value = rv
		}
		return rv, err
	}
	return 0, nil
}

func (v *Value) Float64() float64 {
	rv, _ := v.Float64E()
	return rv
}

func (v *Value) Float32E() (float32, error) {
	rv, err := v.Float64E()
	return float32(rv), err
}

func (v *Value) Float32() float32 {
	return float32(v.Float64())
}

func (v *Value) Uint64E() (uint64, error) {
	if v.value != nil {
		rv, ok := v.value.(uint64)
		if ok {
			return rv, nil
		}
	}
	if v.as == framework.TYPE_UINT {
		rv, err := strconv.ParseUint(v.strValue, 10, 64)
		if err == nil {
			v.value = rv
		}
		return rv, err
	}
	return 0, nil
}

func (v *Value) Uint64() uint64 {
	rv, _ := v.Uint64E()
	return rv
}

func (v *Value) Uint32E() (uint32, error) {
	rv, err := v.Uint64E()
	return uint32(rv), err
}

func (v *Value) Uint32() uint32 {
	return uint32(v.Uint64())
}

func (v *Value) Uint16E() (uint16, error) {
	rv, err := v.Uint64E()
	return uint16(rv), err
}

func (v *Value) Uint16() uint16 {
	return uint16(v.Uint64())
}

func (v *Value) Uint8E() (uint8, error) {
	rv, err := v.Uint64E()
	return uint8(rv), err
}

func (v *Value) Uint8() uint8 {
	return uint8(v.Uint64())
}

func (v *Value) UintE() (uint, error) {
	rv, err := v.Uint64E()
	return uint(rv), err
}

func (v *Value) Uint() uint {
	return uint(v.Uint64())
}
