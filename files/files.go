package files

import (
	"io/ioutil"
	"os"
)

// ファイルの存在確認
func IsExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}
	return nil
}

// 保存
func Save(data []byte, savePath string) error {
	f, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)
	return nil
}

// 読み込み
func Read(readPath string) ([]byte, error) {
	if err := IsExists(readPath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(readPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
