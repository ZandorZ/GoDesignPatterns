package composition

import "testing"

func TestSwimmer_Swim(t *testing.T) {

	swimmer := CompositeSwimmerA{
		MySwim: &Swim,
	}

	swimmer.MyAthlte.Train()
	(*swimmer.MySwim)()

}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()
}

func TestSwimmer_Swim2(t *testing.T) {

	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}

	swimmer.Train()
	swimmer.Swim()

}
