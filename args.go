package nzflag

type NormalizedArgv []Value

// Flags returns only flags in argv
func (argv NormalizedArgv) Flags() []*Flag {
	// flags := make(*Flag, 0, len(values))
	flags := make([]*Flag, 0)
	for _, v := range argv {
		if v.Type() == TypeFlag {
			flags = append(flags, v.Flag())
		}
	}
	return flags
}

// MergedFlags returns flags merged by name
func (argv NormalizedArgv) MergedFlags() []*Flag {
	valuesMap := map[string][]string{}
	indexMap := map[string]int{}
	i := 0
	count := 0
	for _, v := range argv {
		flag := v.Flag()
		if flag == nil {
			continue
		}
		name := flag.Name
		values, ok := valuesMap[name]
		if ok {
			valuesMap[name] = append(values, flag.Values...)
		} else {
			indexMap[name] = i
			valuesMap[name] = flag.Values
			i++
			count++
		}
	}
	flags := make([]*Flag, count)
	for name, values := range valuesMap {
		flags[indexMap[name]] = NewFlag(name, values...)
	}
	return flags
}

// Flags returns only args in argv
func (argv NormalizedArgv) Args() []*Arg {
	// flags := make(*Flag, 0, len(values))
	args := make([]*Arg, 0)
	for _, v := range argv {
		if v.Type() == TypeArg {
			args = append(args, v.Arg())
		}
	}
	return args
}

// Strings returns argv as string slice
func (argv NormalizedArgv) Strings() []string {
	ss := make([]string, len(argv))
	for i, v := range argv {
		ss[i] = v.String()
	}
	return ss
}
