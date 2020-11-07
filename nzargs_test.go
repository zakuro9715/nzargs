package nzargs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleExpected = []string{
	"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x,x",
	"--values1=v", "--values2=v1,v2",
}
var exampleInput = []string{
	"-ab", "-cd=c", "--cd=c", "-ef", "x", "x",
	"--values1=v", "--values2", "v1", "v2",
}

func TestNormalizeToString(t *testing.T) {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	got, err := app.NormalizeToString(exampleInput)
	if assert.NoError(t, err) {
		assert.Equal(t, got, exampleExpected)
	}
}

func ExampleApp_NormalizeToString() {
	app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)
	out, err := app.NormalizeToString(exampleInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", out)
}
