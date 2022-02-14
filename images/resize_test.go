package images

import (
	"testing"
)

func TestResize(t *testing.T) {
	// テスト用画像作成
	opt := createCreateOption(50, 100)
	err := Create(&opt)
	if err != nil {
		t.Errorf("%v", err)
	}
	img, err := Open(opt.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	// リサイズ
	size := Size{
		X: 50,
		Y: 50,
	}
	resizedImg, err := Resize(img, &size, &ResizeOption{
		Fit: Fill,
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	// サイズ確認
	bounds := resizedImg.Bounds()
	if bounds.Dx() != size.X || bounds.Dy() != size.Y {
		t.Error("resize missing")
	}

	// リサイズ済み画像保存
	err = Save(resizedImg, &SaveOption{
		Path:   "sample2.png",
		Format: Png,
	})
	if err != nil {
		t.Errorf("save error: %v", err)
	}
}

func TestResizeContain(t *testing.T) {
	// テスト用画像作成
	opt := createCreateOption(50, 100)
	err := Create(&opt)
	if err != nil {
		t.Errorf("%v", err)
	}
	img, err := Open(opt.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	// リサイズ
	size := Size{
		X: 50,
		Y: 50,
	}
	resizedImg, err := Resize(img, &size, &ResizeOption{
		Fit: Contain,
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	// サイズ確認
	bounds := resizedImg.Bounds()
	if bounds.Dx() != 25 || bounds.Dy() != 50 {
		t.Error("resize missing")
	}

	// リサイズ済み画像保存
	err = Save(resizedImg, &SaveOption{
		Path:   "sample_contain.png",
		Format: Png,
	})
	if err != nil {
		t.Errorf("save error: %v", err)
	}

	// リサイズ
	{
		size := Size{
			X: 25,
			Y: 25,
		}
		resizedImg, err := Resize(img, &size, &ResizeOption{
			Fit: Contain,
		})
		if err != nil {
			t.Errorf("%v", err)
		}

		// サイズ確認
		bounds := resizedImg.Bounds()
		if bounds.Dx() != 12 || bounds.Dy() != 25 {
			t.Error("resize missing")
		}

		// リサイズ済み画像保存
		err = Save(resizedImg, &SaveOption{
			Path:   "sample_contain2.png",
			Format: Png,
		})
		if err != nil {
			t.Errorf("save error: %v", err)
		}

	}
}
