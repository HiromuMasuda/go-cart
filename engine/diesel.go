package engine

import (
	"fmt"
)

type Diesel struct {
}

func NewDiesel() *Diesel {
	return &Diesel{}
}

func (diesel *Diesel) Start() {
	fmt.Println("diesel engine start")
}

func (diesel *Diesel) Stop() {
	fmt.Println("diesel engine stop")
}
