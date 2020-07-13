package abstract

import (
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (c *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SporteMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}
