package prototype

import "testing"

func TestClone(t *testing.T) {

	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Not implemented yet")
	}

	_, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

}
