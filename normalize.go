package nzargv

import (
	"os"
	"strings"
)

// App holds config for normalize
type App struct {
	FlagsValueN map[string]int
}

// New returns new App instance
func New() *App {
	return &App{map[string]int{}}
}

// FlagMaxN sets count of flag value
func (app *App) FlagMaxN(name string, n int) *App {
	app.FlagsValueN[name] = n
	return app
}

// GetFlagMaxN returns count of flag value
func (app *App) GetFlagMaxN(name string) int {
	n, ok := app.FlagsValueN[name]
	if ok {
		return n
	}
	return 0
}

func readFlagValues(max int, args []string) ([]string, int) {
	i := 0
	for ; i < len(args) && i < max && !strings.HasPrefix(args[i], "-"); i++ {
	}
	return args[:i], i
}

func splitByEq(s string) (string, string) {
	splited := strings.SplitN(s, "=", 2)
	if len(splited) == 1 {
		return splited[0], ""
	}
	return splited[0], splited[1]
}

func (app *App) processLongFlag(prefix string, args []string) (Value, int, error) {
	i := 0
	text := strings.TrimPrefix(args[i], prefix)
	name, value := splitByEq(text)
	var flag *Flag
	if len(value) == 0 {
		max := app.GetFlagMaxN(name)
		values, n := readFlagValues(max, args[1:])
		flag = NewFlag(name, values...)
		i += n
	} else {
		values := strings.Split(value, ",")
		flag = NewFlag(name, values...)
	}
	return flag, i, nil
}

func (app *App) processShortFlag(prefix string, args []string) ([]Value, int, error) {
	text := strings.TrimPrefix(args[0], prefix)
	i := 0
	names, value := splitByEq(text)
	lastName := string(names[len(names)-1])
	flags := make([]Value, 0, len(names))

	for _, name := range names[:len(names)-1] {
		flags = append(flags, NewFlag(string(name)))
	}

	if len(value) == 0 {
		max := app.GetFlagMaxN(lastName)
		values, n := readFlagValues(max, args[1:])
		flags = append(flags, NewFlag(lastName, values...))
		i += n
	} else {
		values := strings.Split(value, ",")
		flags = append(flags, NewFlag(lastName, values...))
	}
	return flags, i, nil
}

// Normalize parses argv
func (app *App) Normalize(argv []string) (NormalizedArgv, error) {
	normalized := make([]Value, 0)
	forceArgMode := false
	for i := 0; i < len(argv); i++ {
		v := argv[i]
		switch {
		case v == "--":
			forceArgMode = true
		case forceArgMode:
			normalized = append(normalized, NewArg(v))
		case len(strings.Trim(v, "-")) == 0: // hyphen only
			normalized = append(normalized, NewArg(v))
		case strings.HasPrefix(v, "--"):
			flag, n, err := app.processLongFlag("--", argv[i:])
			if err != nil {
				return nil, err
			}
			i += n
			normalized = append(normalized, flag)

		case strings.HasPrefix(v, "-"):
			flags, n, err := app.processShortFlag("-", argv[i:])
			if err != nil {
				return nil, err
			}
			i += n
			normalized = append(normalized, flags...)
		default:
			normalized = append(normalized, NewArg(v))
		}
	}
	return normalized, nil
}

// NormalizeToStrings normalize argv and returns result as text slice
func (app *App) NormalizeToStrings(argv []string) ([]string, error) {
	normalized, err := app.Normalize(argv)
	if err != nil {
		return nil, err
	}
	return normalized.Strings(), nil
}

// NormalizeArgs is same Normalize except use os.Args
func (app *App) NormalizeArgs() ([]Value, error) {
	return app.Normalize(os.Args[1:])
}

// NormalizeArgsToStrings is same NormalizeToStrings except use os.Args
func (app *App) NormalizeArgsToStrings() ([]string, error) {
	return app.NormalizeToStrings(os.Args[1:])
}
