package pathes

import (
	"os"
	"path"
	"strings"
)

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
func SafeFilename(filename string) string {
	f := filename
	f = strings.ReplaceAll(f, ":", ".")
	f = strings.ReplaceAll(f, "/", "-")
	f = strings.ReplaceAll(f, "|", "-")
	return f
}
