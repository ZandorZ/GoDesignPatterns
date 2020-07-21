package observer

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (s *Publisher) AddObserver(observer ...Observer) {
	s.ObserverList = append(s.ObserverList, observer...)
}

func (s *Publisher) RemoveObserver(observer Observer) {

	var indexRemove int

	for i, ob := range s.ObserverList {
		if ob == observer {
			indexRemove = i
			println("found")
		}
	}

	s.ObserverList = append(s.ObserverList[:indexRemove], s.ObserverList[indexRemove+1:]...)
}

func (s *Publisher) NotifyObservers(m string) {

	for _, ob := range s.ObserverList {
		ob.Notify(m)
	}

}
