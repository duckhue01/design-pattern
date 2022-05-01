package singleton

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

var inst *singleton

func GetInst() Singleton {
	if inst == nil {
		inst = new(singleton)
		inst.AddOne()
	}
	
	return inst
}
