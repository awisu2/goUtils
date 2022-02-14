package images

import (
	"image"

	"golang.org/x/image/draw"
)

// 画像のリサイズ
func Resize(img image.Image, size *Size, opt *ResizeOption) (*image.RGBA, error) {
	// 変換後サイズ計算
	if opt.Fit == Contain || opt.Fit == Cover {
		imgSize := CreateSize(&img)

		rate := GetResizeRate(&imgSize, size, opt.Fit)
		size.X = int(float32(imgSize.X) * rate)
		size.Y = int(float32(imgSize.Y) * rate)
	}

	// 変換後画像オブジェクト作成
	resizedImg := CreateImage(size)
	bounds := resizedImg.Bounds()

	// もと画像を拡大/縮小して書き込み
	// todo: Kernelを切り替えられるようにしても良い
	// 色が先に塗られていた場合dstが優先される
	// 第5引数の draw.Op: alphaが効く形式のときに有効
	//   draw.Over: 上書き(色は交じる)
	//   draw.Src: dstが優先され色の合成はされない
	draw.CatmullRom.Scale(resizedImg, bounds, img, img.Bounds(), draw.Src, nil)
	return resizedImg, nil
}

type ResizeOption struct {
	Fit Fit
}

// get Size from image
func CreateSize(img *image.Image) Size {
	return Size((*img).Bounds().Size())
}

// get AspectRatio
//
// if denominatorY == true: x/y else y/x
//
func GetAspectRatio(img *image.Image, denominatorY bool) float32 {
	size := CreateSize(img)
	return size.AspectRatio(denominatorY)
}

// get resizeRate
func GetResizeRate(srcSize *Size, distSize *Size, fit Fit) float32 {
	srcRatio := srcSize.AspectRatio(true)
	distRatio := distSize.AspectRatio(true)

	var rate float32 = 0
	switch fit {
	case Contain:
		if distRatio >= srcRatio {
			rate = float32(distSize.Y) / float32(srcSize.Y)
		} else {
			rate = float32(distSize.X) / float32(srcSize.X)
		}

	case Cover:
	default:
		if distRatio >= srcRatio {
			rate = float32(distSize.X) / float32(srcSize.X)
		} else {
			rate = float32(distSize.Y) / float32(srcSize.Y)
		}
	}
	return rate
}
