package composition

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
