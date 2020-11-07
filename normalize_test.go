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
		"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x,x",
		"abc", "--values1=v", "--values2=v1,v2", "arg",
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
		NewFlag("f", "x", "x"),
		NewArg("abc"),
		NewFlag("values1", "v"),
		NewFlag("values2", "v1", "v2"),
		NewArg("arg"),
		NewArg("-"),
		NewArg("-----"),
		NewFlag("-v"),
		NewArg("-v=v"),
		NewArg("--value=v"),
	}

	exampleInput = []string{
		"-ab", "-cd=c", "--cd=c", "-ef", "x", "x",
		"abc", "--values1=v", "--values2", "v1", "v2", "arg",
		"-", "-----", "---v",
		"--", "-v=v", "--value=v",
	}
)

func normalizeExampleToStrings() []string {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	ss, err := app.NormalizeToStrings(exampleInput)
	if err != nil {
		panic(err)
	}
	return ss
}

func normalizeExample() NormalizedArgv {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	v, err := app.Normalize(exampleInput)
	if err != nil {
		panic(err)
	}
	return v
}

func ExampleApp_NormalizeToStrings() {
	parsed := normalizeExampleToStrings()
	fmt.Println(strings.Join(parsed[:7], " "))
	fmt.Println(strings.Join(parsed[7:11], " "))
	fmt.Println(strings.Join(parsed[11:], " "))
	// Output:
	// -a -b -c -d=c --cd=c -e -f=x,x
	// abc --values1=v --values2=v1,v2 arg
	// - ----- ---v -v=v --value=v
}

func BenchmarkNormalizeExample(b *testing.B) {
	normalizeExample()
}

func BenchmarkNormalizeExampleToStrings(b *testing.B) {
	normalizeExampleToStrings()
}

func TestNormalize(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	got, err := app.Normalize(exampleInput)
	require.NoError(t, err)
	assert.Equal(t, got, exampleExpectedArgv)
}

func TestNormalizeToStrings(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	got, err := app.NormalizeToStrings(exampleInput)
	require.NoError(t, err)
	assert.Equal(t, got, exampleExpectedStrings)
}

func TestTooFewValues(t *testing.T) {
	app := New().FlagN("value", 2).FlagN("v", 2)
	_, err := app.NormalizeToStrings([]string{"--value", "0"})
	assert.Error(t, err)
	_, err = app.NormalizeToStrings([]string{"-v", "0"})
	assert.Error(t, err)
}

func TestNormalizeArgs(t *testing.T) {
	os.Args = []string{"a", "b", "-c=0"}
	want := []Value{&Arg{"b"}, &Flag{"c", []string{"0"}}}
	got, err := New().NormalizeArgs()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestNormalizeArgsToStrings(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	os.Args = append([]string{"a.out"}, exampleInput...)
	got, err := app.NormalizeArgsToStrings()
	require.NoError(t, err)
	assert.Equal(t, got, exampleExpectedStrings)
}
