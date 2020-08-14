package runtime

import (
	"errors"
	"fmt"
	"github.com/kumpmati/cscr/pkg/ast"
)

// runtime struct
type R struct {
	config  Config
	program *ast.Node
	memory  Memory
}

// returns a new runtime instance
func New() R { return R{} }

// initialize runtime with given config
func (r *R) Init(cfg Config) (err error) {
	r.config = cfg

	r.memory = Memory{}
	r.memory = r.memory.New()
	r.memory.Init()
	return
}

// run the code stored in the program field
func (r *R) Run() (err error) {
	if r.program == nil {
		return errors.New("no program specified")
	}
	if r.config.evaluator == nil {
		return errors.New("no evaluator specified")
	}

	// start evaluating from root node
	result := r.config.evaluator(*r.program, r)
	fmt.Println(result.Get())
	return
}

// set the starting node of the program
func (r *R) SetProgram(node *ast.Node) {
	r.program = node
}
