package abstract

type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

type Car interface {
	GetDoors() int
}

type Motorbike interface {
	GetType() int
}
