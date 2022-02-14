package images

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// 保存オプション
type SaveOption struct {
	Path    string
	Format  Format
	Quality int // quality 1~100 (jpeg用)
}

type CreateOption struct {
	Size
	SaveOption
	Color color.Color
}

func CreateImage(size *Size) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	return img
}

// 画像作成
func Create(option *CreateOption) error {
	img := CreateImage(&option.Size)
	SetRgbasFull(img, option.Color)

	err := Save(img, &option.SaveOption)
	if err != nil {
		return err
	}

	return nil
}

// save to file
//
// if not set format in option, auto ditect from path's ext.
//
func Save(img image.Image, option *SaveOption) (err error) {
	f, err := os.Create(option.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	// auto ditect from path's ext. if not set
	format := option.Format
	if format == "" {
		format = GetFormatFromPath(option.Path)
	}

	switch format {
	case Jpg:
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: option.Quality})
	case Gif:
		err = gif.Encode(f, img, &gif.Options{})
	// case Webp: // Info: go can't create webp(official)
	// case Png:
	default:
		err = png.Encode(f, img)
	}
	if err != nil {
		return err
	}

	return nil
}

// read and create image object
func Open(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// drow color to specific area
func SetRgbas(img *image.RGBA, dr image.Rectangle, color color.Color) {
	for y := dr.Min.Y; y < dr.Max.Y; y++ {
		for x := dr.Min.X; x < dr.Max.X; x++ {
			img.Set(x, y, color)
		}
	}
}

// set collor all area
func SetRgbasFull(img *image.RGBA, color color.Color) {
	SetRgbas(img, img.Bounds(), color)
}
