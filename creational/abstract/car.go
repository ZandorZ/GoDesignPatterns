package abstract

type FamilyCar struct{}

func (l *FamilyCar) GetDoors() int {
	return 5
}

func (l *FamilyCar) GetWheels() int {
	return 4
}
func (l *FamilyCar) GetSeats() int {
	return 5
}

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}

func (l *LuxuryCar) GetWheels() int {
	return 4
}
func (l *LuxuryCar) GetSeats() int {
	return 5
}
