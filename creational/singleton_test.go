package creational

import (
	"testing"
)

func TestGeInstance(t *testing.T) {

	counter1 := GeInstance()

	if counter1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nill")
	}

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("After calling for the first time time to count, the count must be 1 but it is :%d\n", currentCount)
	}

	expectedCounter := counter1

	counter2 := GeInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it is different")
	}

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		t.Errorf("After calling 'AddOne using the second counter, the current count must be 2 but was %d\n'", currentCount)
	}

}
