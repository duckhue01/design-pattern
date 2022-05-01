package builder

type (
	BuildProcess interface {
		SetWheels() BuildProcess
		SetSeats() BuildProcess
		SetStructure() BuildProcess
		GetVehicle() VehicleProduct
	}

	ManufacturingDirector struct {
		builder BuildProcess
	}

	VehicleProduct struct {
		Wheels    int
		Seats     int
		Structure string
	}

	CarBuilder struct {
		vehicle VehicleProduct
	}
	BikeBuilder struct{}
)

// ManufacturingDirector
func (md *ManufacturingDirector) Construct() {
	md.builder.SetSeats().SetStructure().SetWheels()
}

func (md *ManufacturingDirector) SetBuilder(b BuildProcess) {
	md.builder = b
}

// CarBuilder
func (cb *CarBuilder) SetWheels() BuildProcess {
	cb.vehicle.Wheels = 4
	return cb

}

func (cb *CarBuilder) SetSeats() BuildProcess {
	cb.vehicle.Seats = 5
	return cb

}

func (cb *CarBuilder) SetStructure() BuildProcess {
	cb.vehicle.Structure = "car"
	return cb
}

func (cb *CarBuilder) GetVehicle() VehicleProduct {
	return cb.vehicle
}

// BikeBuilder
func (bb *BikeBuilder) SetWheels() BuildProcess {
	panic("not implemented") // TODO: Implement
}

func (bb *BikeBuilder) SetSeats() BuildProcess {
	panic("not implemented") // TODO: Implement
}

func (bb *BikeBuilder) SetStructure() BuildProcess {
	panic("not implemented") // TODO: Implement
}

func (bb *BikeBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}
