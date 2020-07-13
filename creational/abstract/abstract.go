package abstract

import (
	"fmt"
)

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {

	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	case CarFactoryType:
		return new(CarFactory), nil
	default:
		return nil, fmt.Errorf("Factory with id %d not recognized", f)

	}
}
