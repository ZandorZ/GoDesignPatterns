package mediator

import "testing"

func Test_Station(t *testing.T) {
	stationManager := newStationManger()
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}
