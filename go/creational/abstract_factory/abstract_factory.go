package abstract_factory

import "fmt"

// we must retrive a Vehicle object using a factory returned by the abstract factory
// the vehicle must be a concrete implementation of a Motorbike or a Car that implement
//  both interfaces

type (
	VehicleFactory interface {
		Build(v int) (Vehicle, error)
	}
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)

	}
}
