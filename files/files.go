package files

import (
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
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

// 特定のディレクトリ内のファイルから一致するものを検索
func FindFiles(dirname string, re *regexp.Regexp) (matches []fs.FileInfo, err error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	matches = []fs.FileInfo{}
	for i, _ := range fileInfos {
		fileinfo := fileInfos[i]
		if fileinfo.IsDir() {
			continue
		}

		if re.MatchString(fileinfo.Name()) {
			matches = append(matches, fileinfo)
		}
	}

	return matches, nil
}

// 特定のディレクトリ内のファイルから一致した最初のものを返却
func FirstFile(dirname string, re *regexp.Regexp) (match fs.FileInfo, err error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for i, _ := range fileInfos {
		fileinfo := fileInfos[i]
		if fileinfo.IsDir() {
			continue
		}

		if re.MatchString(fileinfo.Name()) {
			return fileinfo, err
		}
	}

	return nil, nil
}
