package abstract

import "testing"

func TestMotorBikeTest(t *testing.T) {

	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d", sportBike.GetType())

}
