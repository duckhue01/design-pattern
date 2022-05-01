package abstract_factory

import (
	"reflect"
	"testing"
)

func TestBuildFactory(t *testing.T) {
	type args struct {
		f int
	}
	tests := []struct {
		name    string
		args    args
		want    VehicleFactory
		wantErr bool
	}{
		{
			name: "test case #1: test car factory",
			args: args{
				f: CarFactoryType,
			},
			want:    &CarFactory{},
			wantErr: false,
		}, {
			name: "test case #2: test motorbike factory",
			args: args{
				f: MotorbikeFactoryType,
			},
			want:    &MotorbikeFactory{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildFactory(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
