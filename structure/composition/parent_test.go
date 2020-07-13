package composition

import "testing"

func TestSon_GetParentField(t *testing.T) {
	son := Son{}
	GetParentField(&son.P)
}
