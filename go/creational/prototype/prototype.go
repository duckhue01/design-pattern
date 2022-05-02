package prototype

import (
	"errors"
	"fmt"
)

// acceptace criteria
// to have a shirt-cloner object and interface to ask for different type of shirt
// when you ask for a white shirt, a clone of the white shirt must be made, and the new instance mist be differrent from the original one
// the Stock Keeping Unit (SKU) of the create object should not affect new object creation
// an info method must give me all the information available on the instance fields, including the updated SKU

type (
	ShirtCloner interface {
		// todo: complete meth
		GetClone(s int) (ItemInfoGetter, error)
	}

	ShirtCache struct{}

	ItemInfoGetter interface {
		GetInfo() string
	}

	ShirtColor byte

	Shirt struct {
		Price float32
		SKU   string
		Color ShirtColor
	}
)

const (
	White = 1
	Black = 2
	Blue  = 3
)

var (
	whitePrototype *Shirt = &Shirt{
		Price: 15.00,
		SKU:   "empty",
		Color: White,
	}
	blackPrototype *Shirt = &Shirt{
		Price: 16.00,
		SKU:   "empty",
		Color: Black,
	}
	bluePrototype *Shirt = &Shirt{
		Price: 17.00,
		SKU:   "empty",
		Color: Blue,
	}
)

func GetShirtCloner() ShirtCloner {
	return &ShirtCache{}
}

func (sc *ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs%f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
