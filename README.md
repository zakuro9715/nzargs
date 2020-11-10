# nzargv


![Test](https://github.com/zakuro9715/nzargv/workflows/Test/badge.svg)
[![codecov](https://codecov.io/gh/zakuro9715/nzargv/branch/main/graph/badge.svg?token=K937ZYFF9Z)](https://codecov.io/gh/zakuro9715/nzargv)
[![GoDoc](https://godoc.org/github.com/zakuro9715/nzargv?status.svg)](http://godoc.org/github.com/zakuro9715/nzargv)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/zakuro9715/nzargv)](https://pkg.go.dev/github.com/zakuro9715/nzargv)
[![Go Report Card](https://goreportcard.com/badge/github.com/zakuro9715/nzargv)](https://goreportcard.com/report/github.com/zakuro9715/nzargv)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Go library to Normalize argv

From
```
-ab=0 --value1=0 --value2 1
```

To
```
-a -b=0 --value=0 --value2=1
```


# Install

```
go get github.com/zakuro9715/nzargv
```

# Motivation

There are useful and powerful cli libraries for Go.

But I want preprocess argv before run cli library and raw cli args are complex format.
Therefore I created nzargv to preprocess raw argv.

# Concept

- Simple
  - No validation. No typecheck
- Respect raw argv. So don't swap order as possible as.
- Result of NormalizeToStrings() is valid argv. So it comopatible with other cli Library (like [urfave/cli](https://github.com/urfave/cli]).
   - You jest pass result of NormalizeToStrings() as Argv

# Feature

- Split short flags (-ab into -a -b)
- Normalize flag with value (--value 0 into value=0)
- Treat as arg after "--" (-f2 -- -f2 # -f2 is arg)
- Merge flags specified multiple times (-a=0 -a=1 into -a=0,1)

# Usage

```input
app := New().FlagN("values1", 2).FlagN("values2", 2).FlagN("f", 2)


app.NormalizeToStrings([]string{
	"-ab", "-cd=c", "--cd=c", "-ef", "x", "x",
   	"--values1=v", "--values2", "v1", "v2", "arg",
});

// Result
[]string{
	"-a", "-b", "-c", "-d=c", "--cd=c", "-e", "-f=x,x",
   	"--values1=v", "--values2=v1,v2", "arg"
}
```

See also ![PkgGoDev](https://pkg.go.dev/badge/github.com/zakuro9715/nzargv)](https://pkg.go.dev/github.com/zakuro9715/nzargv) and [test](normalize_test.go)

# CLI Tool

## Install

```
go get github.com/zakuro9715/nzargv/cmd/nzargv
```
