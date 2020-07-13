package composition

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

type SwimClosureType func()

var Swim SwimClosureType = func() {
	println("Swimming")
}

type CompositeSwimmerA struct {
	MyAthlte Athlete
	MySwim   *SwimClosureType
}

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// -------------------------------------------------

type Animal struct{}

func (a *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim SwimClosureType
}
