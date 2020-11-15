package nzflag

import (
	"fmt"
	"strings"
)

// ValueType represents type of Value
type ValueType int

const (
	// TypeFlag is flag
	TypeFlag ValueType = iota
	// TypeArg is arg
	TypeArg
)

// Value represents Arg or Flag
type Value interface {
	fmt.Stringer
	Type() ValueType
	Flag() *Flag
	Arg() *Arg
}

// Flag is cli flag
type Flag struct {
	Name   string
	Values []string
}

// NewFlag returns flag instance
func NewFlag(name string, values ...string) *Flag {
	if values == nil {
		values = []string{}
	}
	return &Flag{name, values}
}

// Flag returns itself
func (v *Flag) Flag() *Flag {
	return v
}

// Arg returns nil
func (v *Flag) Arg() *Arg {
	return nil
}

// Type returns TypeFlag
func (v *Flag) Type() ValueType {
	return TypeFlag
}

// String returns flag string
func (v *Flag) String() string {
	name := v.Name
	if len(name) == 1 {
		name = "-" + v.Name
	} else {
		name = "--" + v.Name
	}
	value := strings.Join(v.Values, ",")
	if len(value) == 0 {
		return name
	}
	return fmt.Sprintf("%v=%v", name, value)
}

// Arg represets cli argument
type Arg struct {
	Value string
}

// NewArg returns new Arg instance
func NewArg(value string) *Arg {
	return &Arg{value}
}

// Flag returns nil
func (v *Arg) Flag() *Flag {
	return nil
}

// Arg returns itself
func (v *Arg) Arg() *Arg {
	return v
}

// Type returns TypeArg
func (v *Arg) Type() ValueType {
	return TypeArg
}

// String returns arg string
func (v *Arg) String() string {
	return v.Value
}
