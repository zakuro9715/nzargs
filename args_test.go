package nzargv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagsAndArgs(t *testing.T) {
	argv := normalizeExample()
	flags := []*Flag{
		NewFlag("a"),
		NewFlag("b"),
		NewFlag("c"),
		NewFlag("d", "c"),
		NewFlag("cd", "c"),
		NewFlag("e"),
		NewFlag("f", "x", "x"),
		NewFlag("values1", "v"),
		NewFlag("values2", "v1", "v2"),
	}
	args := []*Arg{
		NewArg("abc"),
		NewArg("arg"),
	}
	assert.Equal(t, flags, argv.Flags())
	assert.Equal(t, args, argv.Args())
}

func TestMergedFlags(t *testing.T) {
	assert.Equal(t,
		[]*Flag{
			NewFlag("a", "a"),
			NewFlag("b"),
			NewFlag("values", "1", "2", "1", "3"),
			NewFlag("c"),
			NewFlag("d"),
		},
		NormalizedArgv{
			NewFlag("a", "a"),
			NewFlag("b"),
			NewFlag("values", "1", "2"),
			NewFlag("c"),
			NewArg("dummy"),
			NewFlag("values", "1", "3"),
			NewFlag("d"),
		}.MergedFlags(),
	)
}
