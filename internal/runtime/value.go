package runtime

// a value must have a function that returns its value
type Value interface {
	Get() float32
}

type Constant struct {
	value float32
}

type Reference struct {
	ref *Variable
}

func (c Constant) Get() float32 {
	return c.value
}

// a reference's value is the value of its reference.
func (r Reference) Get() float32 {
	return r.ref.value.Get()
}
