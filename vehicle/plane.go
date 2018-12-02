package vehicle

import (
	"../engine"
	"fmt"
)

type Plane struct {
	Engine engine.Engine
}

func NewPlane(engine engine.Engine) *Plane {
	return &Plane{
		Engine: engine,
	}
}

func (plane *Plane) Travel() {
	plane.Engine.Start()
	fmt.Println("Plane travel")
	plane.Engine.Stop()
}
