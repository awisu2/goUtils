package images

import (
	"testing"
)

func createCreateOption(x int, y int) CreateOption {
	return CreateOption{
		Size: Size{
			X: x,
			Y: y,
		},
		SaveOption: SaveOption{
			Path:    "sample.png",
			Format:  Png,
			Quality: 70,
		},
		Color: RgbaGreen,
	}

}

func TestCreate(t *testing.T) {
	opt := createCreateOption(50, 100)
	err := Create(&opt)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestOpen(t *testing.T) {
	opt := createCreateOption(50, 100)
	err := Create(&opt)
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = Open(opt.Path)
	if err != nil {
		t.Errorf("%v", err)
	}
}
