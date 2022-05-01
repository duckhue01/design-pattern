package abstract_factory

type (
	Motorbike interface {
		GetMotorbikeType() int
	}
	SportMotorbike struct {
	}

	CruiseMotorbike struct {
	}
)

func (sm *SportMotorbike) NumWheels() int {
	panic("not implemented") // TODO: Implement
}

func (sm *SportMotorbike) NumSeats() int {
	panic("not implemented") // TODO: Implement
}

func (cm *CruiseMotorbike) NumWheels() int {
	panic("not implemented") // TODO: Implement
}

func (cm *CruiseMotorbike) NumSeats() int {
	panic("not implemented") // TODO: Implement
}
