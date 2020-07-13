package prototype

import (
	"fmt"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCache struct {
	clones map[int]*Shirt
}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	if _, ok := s.clones[m]; !ok {
		return nil, fmt.Errorf("Shirt model %d not recognized", m)
	} else {
		clone := *s.clones[m]
		return &clone, nil
	}

}

type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU %s and Color id %d that costs %f", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{
		clones: map[int]*Shirt{
			White: &Shirt{
				Price: 15.00,
				SKU:   "Empty",
				Color: White,
			},
			Black: &Shirt{
				Price: 16.00,
				SKU:   "Empty",
				Color: Black,
			},
			Blue: &Shirt{
				Price: 17.00,
				SKU:   "Empty",
				Color: Blue,
			},
		},
	}
}
