package vehicle

import (
	"../engine"
	"fmt"
)

type Car struct {
	Engine engine.Engine
}

func NewCar(engine engine.Engine) *Car {
	return &Car{
		Engine: engine,
	}
}

func (car *Car) Travel() {
	car.Engine.Start()
	fmt.Println("Car travel")
	car.Engine.Stop()
}
