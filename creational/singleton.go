package creational

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

func GeInstance() Singleton {
	if instance == nil {
		instance = &singleton{count: 0}
	}
	return instance
}
