package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message string
}

func (o *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message: '%s' received", o.ID, m)
	o.Message = m
}

func TestSubject(t *testing.T) {
	testOb1 := &TestObserver{1, "default"}
	testOb2 := &TestObserver{2, "default"}
	testOb3 := &TestObserver{3, "default"}

	pub := Publisher{}

	t.Run("Add Observer", func(t *testing.T) {

		pub.AddObserver(testOb1, testOb2, testOb3)

		if len(pub.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("Remove Observer", func(t *testing.T) {
		pub.RemoveObserver(testOb2)

		if len(pub.ObserverList) != 2 {
			t.Errorf("The size of the observer list is not the expected. 3 != %d\n", len(pub.ObserverList))
		}

		for _, observer := range pub.ObserverList {
			testOb, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testOb.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {

		for _, observer := range pub.ObserverList {
			printOb, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printOb.Message != "default" {
				t.Errorf("The observer's Message field weren't empty: %s\n", printOb.Message)
			}
		}

		message := "Hello World!"
		pub.NotifyObservers(message)

		for _, observer := range pub.ObserverList {
			printOb, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printOb.Message != message {
				t.Errorf("Expected message on observer %d was not expected: '%s' != '%s'\n", printOb.ID, printOb.Message, message)
			}
		}

	})

}
