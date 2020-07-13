package composition

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p *Parent) int {
	return p.SomeField
}
