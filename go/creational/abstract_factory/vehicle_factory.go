package abstract_factory

import (
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2

	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type (

	CarFactory struct{}

	MotorbikeFactory struct{}
)

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil

	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}
