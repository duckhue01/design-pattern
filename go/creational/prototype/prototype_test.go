package prototype

import (
	"reflect"
	"testing"
)

func TestGetShirtCloner(t *testing.T) {
	tests := []struct {
		name string
		want ShirtCloner
	}{
		{
			name: "test case #1: ",
			want: &ShirtCache{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetShirtCloner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShirtCloner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShirtCache_GetClone(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name    string
		sc      *ShirtCache
		args    args
		want    ItemInfoGetter
		wantErr bool
	}{
		{
			name: "test case #1: want to get cloned instance",
			sc:   &ShirtCache{},
			args: args{
				s: White,
			},
			want:    whitePrototype,
			wantErr: false,
		},
		{
			name: "test case #2: want to get cloned instance",
			sc:   &ShirtCache{},
			args: args{
				s: Blue,
			},
			want:    bluePrototype,
			wantErr: false,
		},
		{
			name: "test case #3: want to get cloned instance",
			sc:   &ShirtCache{},
			args: args{
				s: Black,
			},
			want:    blackPrototype,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &ShirtCache{}
			got, err := sc.GetClone(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShirtCache.GetClone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShirtCache.GetClone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShirtCache_GetClone_Compare(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name    string
		sc      *ShirtCache
		args    args
		wantErr bool
	}{
		{
			name: "test case #1: want to get cloned instance",
			sc:   &ShirtCache{},
			args: args{
				s: White,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &ShirtCache{}
			got1, err := sc.GetClone(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShirtCache.GetClone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got2, err := sc.GetClone(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShirtCache.GetClone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got1 == got2 {
				t.Errorf("two instance must be allocated on difference location %v , %v", &got1, &got2)
			}

			t.Logf("LOG: %s", got1.GetInfo())
			t.Logf("LOG: %s", got2.GetInfo())
		})
	}
}

func TestShirt_GetInfo(t *testing.T) {
	type fields struct {
		Price float32
		SKU   string
		Color ShirtColor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shirt{
				Price: tt.fields.Price,
				SKU:   tt.fields.SKU,
				Color: tt.fields.Color,
			}
			if got := s.GetInfo(); got != tt.want {
				t.Errorf("Shirt.GetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShirt_GetPrice(t *testing.T) {
	type fields struct {
		Price float32
		SKU   string
		Color ShirtColor
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shirt{
				Price: tt.fields.Price,
				SKU:   tt.fields.SKU,
				Color: tt.fields.Color,
			}
			if got := s.GetPrice(); got != tt.want {
				t.Errorf("Shirt.GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
