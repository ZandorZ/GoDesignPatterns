package abstract

type CruiseMotorbike struct{}

func (l *CruiseMotorbike) GetWheels() int {
	return 2
}
func (l *CruiseMotorbike) GetSeats() int {
	return 2
}

func (l *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}

type SporteMotorbike struct{}

func (l *SporteMotorbike) GetWheels() int {
	return 2
}
func (l *SporteMotorbike) GetSeats() int {
	return 1
}

func (l *SporteMotorbike) GetType() int {
	return SportMotorbikeType
}
