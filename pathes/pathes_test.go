package pathes

import (
	"strings"
	"testing"
)

func TestUserPathes(t *testing.T) {
	pathes, err := UserPathes()
	if err != nil {
		t.Error(err)
	}
	if pathes.HomeDir == "" || pathes.CacheDir == "" || pathes.ConfigDir == "" {
		t.Errorf("not set directory")
	}
}

func TestAppConfigDir(t *testing.T) {
	appName := "testApp"
	appConfigDir, err := AppConfigDir(appName)
	if err != nil {
		t.Error(err)
	}
	if appConfigDir == "" {
		t.Error(err)
	}
	if !strings.Contains(appConfigDir, appName) {
		t.Error(err)
	}
}

func TestAppConfigPath(t *testing.T) {
	appName := "testApp"
	configName := "config"
	appConfigPath, err := AppConfigPath(appName, configName)
	if err != nil {
		t.Error(err)
	}
	if appConfigPath == "" {
		t.Error(err)
	}
	if !strings.Contains(appConfigPath, appName) {
		t.Error(err)
	}
	if !strings.Contains(appConfigPath, configName) {
		t.Error(err)
	}
}
