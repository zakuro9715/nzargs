package nzflag

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	exampleExpectedStrings = []string{
		"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x", "-g=x=x", "abc",
		"--value=v1", "--value=v2", "--value", "--help",
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
		NewFlag("g", "x=x"),
		NewArg("abc"),
		NewFlag("value", "v1"),
		NewFlag("value", "v2"),
		NewFlag("value"),
		NewFlag("help"),
		NewArg("-"),
		NewArg("-----"),
		NewFlag("-v"),
		NewArg("-v=v"),
		NewArg("--value=v"),
	}

	exampleInput = []string{
		"-ab", "-cd=c", "--cd=c", "-ef", "x", "-gx=x", "abc",
		"--value=v1", "--value", "v2", "--value", "--help",
		"-", "-----", "---v",
		"--", "-v=v", "--value=v",
	}
)

func TestAppFlagPanic(t *testing.T) {
	assert.Panics(t, func() { New().Flag("f", -1) })
}

func TestAppFlag(t *testing.T) {
	app := New()
	app.Flag("f", HasValue)
	assert.True(t, app.FlagHasValue("f"))
	app.Flag("f", None)
	assert.Equal(t, None, app.FlagOption["f"])
	app.Flag("fa", None)
	assert.Equal(t, None, app.FlagOption["f"])
}

var exampleApp = &App{
	FlagOption: map[string]FlagOption{
		"f":     HasValue,
		"g":     HasValue,
		"value": HasValue,
	},
}

func normalizeExampleToStrings() []string {
	return exampleApp.NormalizeToStrings(exampleInput)
}

func normalizeExample() NormalizedArgv {
	return exampleApp.Normalize(exampleInput)
}

func ExampleApp_NormalizeToStrings() {
	parsed := normalizeExampleToStrings()
	fmt.Println(strings.Join(parsed[:9], " "))
	fmt.Println(strings.Join(parsed[9:13], " "))
	fmt.Println(strings.Join(parsed[13:], " "))
	// Output:
	// -a -b -c -d=c --cd=c -e -f=x -g=x=x abc
	// --value=v1 --value=v2 --value --help
	// - ----- ---v -v=v --value=v
}

func BenchmarkNormalizeExample(b *testing.B) {
	normalizeExample()
}

func BenchmarkNormalizeExampleToStrings(b *testing.B) {
	normalizeExampleToStrings()
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, exampleExpectedArgv, exampleApp.Normalize(exampleInput))
}

func TestNormalizeToStrings(t *testing.T) {
	assert.Equal(t, exampleExpectedStrings, exampleApp.NormalizeToStrings(exampleInput))
}

func TestNormalizeArgs(t *testing.T) {
	os.Args = []string{"a", "b", "-c=0"}
	want := []Value{&Arg{"b"}, &Flag{"c", []string{"0"}}}
	got := New().NormalizeArgs()
	assert.Equal(t, want, got)
}

func TestNormalizeArgsToStrings(t *testing.T) {
	os.Args = append([]string{"a.out"}, exampleInput...)
	got := exampleApp.NormalizeArgsToStrings()
	assert.Equal(t, exampleExpectedStrings, got)
}
