package runtime

import (
	"fmt"
)

type Memory struct {
	Variables map[string]*Variable
}

type Variable struct {
	name  string
	value Value
}

func (m *Memory) New() Memory {
	return Memory{}
}

func (m *Memory) Init() {
	m.Variables = map[string]*Variable{}
}

func (m *Memory) GetVar(name string) *Variable {
	if _, ok := m.Variables[name]; ok {
		return m.Variables[name]
	}
	return nil
}

func (m *Memory) NewVariable(name string, value Value) *Variable {
	// check that name does not exist
	for _, v := range m.Variables {
		if v.name == name {
			fmt.Printf("'%v' already exists\n", name)
			return nil
		}
	}

	newVar := Variable{name: name, value: value}
	// add variable to map
	m.Variables[name] = &newVar
	return &newVar
}
