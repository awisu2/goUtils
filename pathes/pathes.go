package pathes

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/awisu2/goUtils/files"
)

// windowsでファイル名に拒否される文字列
const windowsIgnoreWords = `<>:"/\|?*`

type GetPathResponse struct {
	ConfigDir string
	HomeDir   string
	CacheDir  string
}

// ユーザ関連ディレクトリの取得
//
// windowsの場合
// C:\Users\{user}\AppData\Roaming <nil>
// C:\Users\{user} <nil>
// C:\Users\{user}\AppData\Local <nil>
func UserPathes() (GetPathResponse, error) {

	userPathes := GetPathResponse{}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return userPathes, err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return userPathes, err
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return userPathes, err
	}

	userPathes.ConfigDir = configDir
	userPathes.HomeDir = homeDir
	userPathes.CacheDir = cacheDir

	return userPathes, nil
}

// userConfig配下のアプリディレクトリパスを生成
func AppConfigDir(appName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appConfigDir := path.Join(configDir, appName)
	return appConfigDir, err
}

// userConfig配下のアプリconfigファイルパスを生成
func AppConfigPath(appName string, configName string) (string, error) {
	appConfigDir, err := AppConfigDir(appName)
	if err != nil {
		return "", err
	}

	appConfigPath := path.Join(appConfigDir, configName)
	return appConfigPath, err
}

// osに対して安全なファイル名に変換して返却
func SafeFilename(filename string, opt SafeFilenameOption) string {
	f := filename
	for _, c := range windowsIgnoreWords {
		f = strings.ReplaceAll(f, string(c), opt.Alias)
	}

	return f
}

type SafeFilenameOption struct {
	Alias string // 置換文字
}

var DefaultSafeFileNameOption = SafeFilenameOption{
	Alias: " ",
}

// make safe path
//
// 1. if already exists file or directory change pathname
func MakeSafePath(path string, opt MakeSafePathOption) string {
	// check file dupulicatte
	if opt.FixDuplicate {
		_path := path

		name, ext := path, ""
		if !opt.IsDirectory {
			name, ext = SplitName(path)
		}
		for i := 1; ; i = i + 1 {
			if files.IsExists(_path) {
				_path = fmt.Sprintf("%s_%d%s", name, i, ext)
			} else {
				break
			}
		}
		path = _path
	}
	return path
}

type MakeSafePathOption struct {
	IsDirectory  bool
	FixDuplicate bool
}

// rename base without extension
func ReBaseName(path string, change string, opt ReBaseNameOption) string {
	dir, file := filepath.Split(path)
	if !opt.IsDirectory {
		ext := filepath.Ext(file)
		file = change + ext
	} else {
		file = change
	}
	return filepath.Join(dir, file)
}

type ReBaseNameOption struct {
	IsDirectory bool // if true overwrite with ext
}

// split file name to name and extension
func SplitName(file string) (name string, ext string) {
	ext = filepath.Ext(file)
	if ext == "" {
		return file, ""
	}
	return file[:len(file)-len(ext)], ext
}
