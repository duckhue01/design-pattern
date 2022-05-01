package singleton

import (
	"reflect"
	"testing"
)

func TestGetInst(t *testing.T) {
	tests := []struct {
		name string
		want Singleton
	}{
		{
			name: "test case #1: doesn't has any instant yet need to create one",
			want: &singleton{
				count: 1,
			},
		},
		{
			name: "test case #2: an instant already exist",
			want: &singleton{
				count: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleton_AddOne(t *testing.T) {
	type fields struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &singleton{
				count: tt.fields.count,
			}
			if got := s.AddOne(); got != tt.want {
				t.Errorf("singleton.AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
