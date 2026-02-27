package simulation

import (
	"errors"

	"github.com/lorenz-xyz/tire-simuation/internal/utilities"
)

type Model struct {
	SurfaceModel   *Surface
	CarModel       *Car
	FrontLeftTire  *Tire
	FrontRightTire *Tire
	RearLeftTire   *Tire
	RearRightTire  *Tire
}

type Surface struct {
	Seed int
}

type Car struct {
	WeightKg         int
	WheelBaseM       float32
	RatioWeightFront float32
	RatioWeightRear  float32
	CenterOfMass     utilities.Vector
}

func NewCar(weight int, wheelbase float32, ratioFront float32, ratioRear float32) (Car, error) {
	if weight < 0 || wheelbase < 0 || ratioFront < 0 || ratioRear < 0 {
		return Car{}, errors.New("No negative numbers possible")
	}
	if ratioFront > 1 || ratioRear > 1 {
		return Car{}, errors.New("Ratios need to be between 0 and 1")
	}

	return Car{
		WeightKg:         weight,
		RatioWeightFront: ratioFront,
		RatioWeightRear:  ratioRear,
	}, nil
}

type Tire struct {
	TireWidthMm     int
	TireHeightRatio int
	WheelDiameter   int
}
