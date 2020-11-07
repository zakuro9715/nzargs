package nzargs

import (
	"fmt"
	"strings"
)

type ValueType int

const (
	TypeFlag ValueType = iota
	TypeArg
)

type Value interface {
	Type() ValueType
	Text() string
	Flag() *Flag
	Arg() *Arg
}

type Flag struct {
	Name   string
	Values []string
}

func NewFlag(name string, values ...string) *Flag {
	return &Flag{name, values}
}

func (v *Flag) Flag() *Flag {
	return v
}

func (v *Flag) Arg() *Arg {
	return nil
}

func (v *Flag) Type() ValueType {
	return TypeFlag
}

func (v *Flag) Text() string {
	name := v.Name
	if len(name) == 1 {
		name = "-" + v.Name
	} else {
		name = "--" + v.Name
	}
	if len(v.Values) == 0 {
		return name
	}
	return fmt.Sprintf("%v=%v", name, strings.Join(v.Values, ","))
}

type Arg struct {
	Value string
}

func NewArg(value string) *Arg {
	return &Arg{value}
}

func (v *Arg) Flag() *Flag {
	return nil
}

func (v *Arg) Arg() *Arg {
	return v
}

func (v *Arg) Type() ValueType {
	return TypeArg
}

func (v *Arg) Text() string {
	return v.Value
}
