package runtime

import "github.com/kumpmati/cscr/pkg/args"

// runtime config struct
type Config struct {
	evaluator EvalFunc
}

// the default runtime config
func DefaultRuntimeConfig(a args.Args) Config {
	c := Config{}
	c.evaluator = DefaultEvalFunc
	return c
}
