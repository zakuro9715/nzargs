package nzargv

import (
	"fmt"
	"os"
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

func ExampleApp_NormalizeToStrings() {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	out, err := app.NormalizeToStrings(exampleInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", out)
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
	_, err := app.Normalize([]string{"--value", "0"})
	assert.Error(t, err)
	_, err = app.Normalize([]string{"-v", "0"})
	assert.Error(t, err)
}

func TestNormalizeArgsToStrings(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	os.Args = append([]string{"a.out"}, exampleInput...)
	got, err := app.NormalizeArgsToStrings()
	if assert.NoError(t, err) {
		assert.Equal(t, got, exampleExpected)
	}
}
