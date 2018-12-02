package engine

import (
	"fmt"
)

type Gasoline struct {
}

func NewGasoline() *Gasoline {
	return &Gasoline{}
}

func (gasoline *Gasoline) Start() {
	fmt.Println("gasoline engine start")
}

func (gasoline *Gasoline) Stop() {
	fmt.Println("gasoline engine stop")
}
