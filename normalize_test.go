package nzargv

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleExpected = []string{
	"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x,x",
	"--values1=v", "--values2=v1,v2", "arg",
}
var exampleInput = []string{
	"-ab", "-cd=c", "--cd=c", "-ef", "x", "x",
	"--values1=v", "--values2", "v1", "v2", "arg",
}

func normalizeExampleToStrings() ([]string, error) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	return app.NormalizeToStrings(exampleInput)
}

func normalizeExample() ([]string, error) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	return app.NormalizeToStrings(exampleInput)
}

func ExampleApp_NormalizeToStrings() {
	parsed, err := normalizeExampleToStrings()
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(parsed[:7], " "))
	fmt.Println(strings.Join(parsed[7:], " "))
	// Output:
	// -a -b -c -d=c --cd=c -e -f=x,x
	// --values1=v --values2=v1,v2 arg
}

func TestNormalizeToStrings(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	got, err := app.NormalizeToStrings(exampleInput)
	if assert.NoError(t, err) {
		assert.Equal(t, got, exampleExpected)
	}
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
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestNormalizeArgsToStrings(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	os.Args = append([]string{"a.out"}, exampleInput...)
	got, err := app.NormalizeArgsToStrings()
	if assert.NoError(t, err) {
		assert.Equal(t, got, exampleExpected)
	}
}
