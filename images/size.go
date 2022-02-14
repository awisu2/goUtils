package images

import "image"

// Info: Point structがimage packageに含まれるが用途が違うため新規宣言
type Size image.Point

// get AspectRatio
//
// if denominatorHeight == true: width/height else height/width
//
func (s *Size) AspectRatio(denominatorY bool) float32 {
	if denominatorY {
		if s.Y == 0 {
			return 0
		}
		return float32(s.X) / float32(s.Y)
	} else {
		if s.X == 0 {
			return 0
		}
		return float32(s.Y) / float32(s.X)
	}

	return 1
}

// todo: constで保管しておく方法
var (
	Hd = Size{
		X: 2560,
		Y: 1440,
	}
	FullHd = Size{
		X: 1920,
		Y: 1000,
	}
	K2 = Size{
		X: 2560,
		Y: 1440,
	}
	K4 = Size{
		X: 3840,
		Y: 2160,
	}
	K8 = Size{
		X: 7680,
		Y: 4320,
	}
)
