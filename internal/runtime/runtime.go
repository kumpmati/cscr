package runtime

type Config struct{}

// runtime struct
type Runtime struct {
	config Config
}

// returns a new runtime instance
func New() Runtime { return Runtime{} }

// set flags for runtime
func (r *Runtime) Init(cfg Config) (err error) {
	return
}
