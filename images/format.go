package images

import "path/filepath"

// formtディレクトリなど作ってサブパッケージに分けたいがimportが複雑になるので同パッケージで管理
type Format string

const (
	Jpg  Format = "jpg"
	Png  Format = "png"
	Gif  Format = "gif"
	Webp Format = "webp"
)

func GetFormatFromExt(ext string) Format {
	format := Png
	switch ext {
	case ".jpg", ".jpeg":
		format = Jpg
	case ".png":
		format = Png
	case ".gif":
		format = Gif
	case ".webp":
		format = Webp
	}
	return format
}

func GetFormatFromPath(path string) Format {
	ext := filepath.Ext(path)
	return GetFormatFromExt(ext)
}
