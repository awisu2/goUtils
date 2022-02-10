package files

import (
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
)

// check file exists
//
// if ignore error hapen it's return false
// because if all error return this function is too deficult to use
//
func IsExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		// if os.IsNotExist(err) {}
		return false
	}
	return true
}

// save file
func Save(data []byte, savePath string) error {
	f, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)
	return nil
}

// read file
func Read(readPath string) ([]byte, error) {
	if !IsExists(readPath) {
		return nil, os.ErrNotExist
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
