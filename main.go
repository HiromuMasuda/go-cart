package main

import (
	"./engine"
	"./vehicle"
)

func main() {
	engine := engine.NewGasoline()
	vehicle := vehicle.NewCar(engine)
	vehicle.Travel()
}
