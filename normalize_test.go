package nzargv

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	exampleExpectedStrings = []string{
		"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x", "--help",
		"abc", "--value=v1", "--value=v2", "arg",
		"-", "-----", "---v",
		"-v=v", "--value=v",
	}

	exampleExpectedArgv = NormalizedArgv{
		NewFlag("a"),
		NewFlag("b"),
		NewFlag("c"),
		NewFlag("d", "c"),
		NewFlag("cd", "c"),
		NewFlag("e"),
		NewFlag("f", "x"),
		NewFlag("help"),
		NewArg("abc"),
		NewFlag("value", "v1"),
		NewFlag("value", "v2"),
		NewArg("arg"),
		NewArg("-"),
		NewArg("-----"),
		NewFlag("-v"),
		NewArg("-v=v"),
		NewArg("--value=v"),
	}

	exampleInput = []string{
		"-ab", "-cd=c", "--cd=c", "-ef", "x", "--help",
		"abc", "--value=v1", "--value", "v2", "arg",
		"-", "-----", "---v",
		"--", "-v=v", "--value=v",
	}
)

var exampleApp = New().Flag("f", HasValue).Flag("value", HasValue)

func normalizeExampleToStrings() []string {
	ss, err := exampleApp.NormalizeToStrings(exampleInput)
	if err != nil {
		panic(err)
	}
	return ss
}

func normalizeExample() NormalizedArgv {
	v, err := exampleApp.Normalize(exampleInput)
	if err != nil {
		panic(err)
	}
	return v
}

func ExampleApp_NormalizeToStrings() {
	parsed := normalizeExampleToStrings()
	fmt.Println(strings.Join(parsed[:8], " "))
	fmt.Println(strings.Join(parsed[8:12], " "))
	fmt.Println(strings.Join(parsed[12:], " "))
	// Output:
	// -a -b -c -d=c --cd=c -e -f=x --help
	// abc --value=v1 --value=v2 arg
	// - ----- ---v -v=v --value=v
}

func BenchmarkNormalizeExample(b *testing.B) {
	normalizeExample()
}

func BenchmarkNormalizeExampleToStrings(b *testing.B) {
	normalizeExampleToStrings()
}

func TestNormalize(t *testing.T) {
	got, err := exampleApp.Normalize(exampleInput)
	require.NoError(t, err)
	assert.Equal(t, exampleExpectedArgv, got)
}

func TestNormalizeToStrings(t *testing.T) {
	got, err := exampleApp.NormalizeToStrings(exampleInput)
	require.NoError(t, err)
	assert.Equal(t, exampleExpectedStrings, got)
}

func TestNormalizeArgs(t *testing.T) {
	os.Args = []string{"a", "b", "-c=0"}
	want := []Value{&Arg{"b"}, &Flag{"c", []string{"0"}}}
	got, err := New().NormalizeArgs()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestNormalizeArgsToStrings(t *testing.T) {
	os.Args = append([]string{"a.out"}, exampleInput...)
	got, err := exampleApp.NormalizeArgsToStrings()
	require.NoError(t, err)
	assert.Equal(t, exampleExpectedStrings, got)
}
