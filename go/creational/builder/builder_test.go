package builder

import (
	"reflect"
	"testing"
)

// 1. i must have a manufacturing type that constructs everything that vehicle needs

// 2. when using a car builder, the VehicleProduct with four wheels, five seats, and a structure defined as Car must be returned

// 3. when using a motorbike builder, the VehicleProduct with two wheels, two seats, and structure defined as Motorbike must be returned

// 4. a VehicleProduct built by any BuildProcess builder must be open to modifications

func TestManufacturingDirector_Construct(t *testing.T) {
	tests := []struct {
		name string
		md   *ManufacturingDirector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &ManufacturingDirector{}
			md.Construct()
		})
	}
}

func TestManufacturingDirector_SetBuilder(t *testing.T) {
	type args struct {
		b BuildProcess
	}
	tests := []struct {
		name string
		md   *ManufacturingDirector
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &ManufacturingDirector{}
			md.SetBuilder(tt.args.b)
		})
	}
}

func TestCarBuilder_SetWheels(t *testing.T) {
	tests := []struct {
		name string
		cb   *CarBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := &CarBuilder{}
			if got := cb.SetWheels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetWheels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_SetSeats(t *testing.T) {
	tests := []struct {
		name string
		cb   *CarBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := &CarBuilder{}
			if got := cb.SetSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_SetStructure(t *testing.T) {
	tests := []struct {
		name string
		cb   *CarBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := &CarBuilder{}
			if got := cb.SetStructure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_GetVehicle(t *testing.T) {
	tests := []struct {
		name string
		cb   *CarBuilder
		want VehicleProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := &CarBuilder{}
			if got := cb.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetWheels(t *testing.T) {
	tests := []struct {
		name string
		bb   *BikeBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &BikeBuilder{}
			if got := bb.SetWheels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetWheels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetSeats(t *testing.T) {
	tests := []struct {
		name string
		bb   *BikeBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &BikeBuilder{}
			if got := bb.SetSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetStructure(t *testing.T) {
	tests := []struct {
		name string
		bb   *BikeBuilder
		want BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &BikeBuilder{}
			if got := bb.SetStructure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_GetVehicle(t *testing.T) {
	tests := []struct {
		name string
		bb   *BikeBuilder
		want VehicleProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &BikeBuilder{}
			if got := bb.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilderPattern(t *testing.T) {

	tests := []struct {
		name    string
		builder BuildProcess
		expect  VehicleProduct
	}{
		{name: "test case #1: test car builder", builder: &CarBuilder{}, expect: VehicleProduct{
			Wheels:    4,
			Seats:     5,
			Structure: "car",
		}},
		{name: "test case #2: test bike builder", builder: &BikeBuilder{}, expect: VehicleProduct{
			Wheels:    2,
			Seats:     1,
			Structure: "bike",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			md := ManufacturingDirector{}
			builder := &CarBuilder{}
			md.SetBuilder(tt.builder)
			md.Construct()

			vehicle := builder.GetVehicle()

			if vehicle.Wheels != tt.expect.Wheels {
				t.Errorf("car.Wheels = %v, but want %v", vehicle.Wheels, tt.expect.Wheels)
			}
			if vehicle.Seats != tt.expect.Seats {
				t.Errorf("car.Seats = %v, but want %v", vehicle.Seats, tt.expect.Seats)
			}

			if vehicle.Structure != tt.expect.Structure {
				t.Errorf("car.Structure = %v, but want %v", vehicle.Structure, tt.expect.Structure)
			}

		})
	}

}
