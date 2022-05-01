package abstract_factory

type (
	Car interface {
		NumDoors() int
	}

	LuxuryCar struct {
	}

	FamilyCar struct {
	}
)

func (luxc *LuxuryCar) NumWheels() int {
	panic("not implemented") // TODO: Implement
}

func (luxc *LuxuryCar) NumSeats() int {
	panic("not implemented") // TODO: Implement
}

func (famc *FamilyCar) NumWheels() int {
	panic("not implemented") // TODO: Implement
}

func (famc *FamilyCar) NumSeats() int {
	panic("not implemented") // TODO: Implement
}
