package pathes

import (
	"os"
	"path/filepath"
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

func TestSafePath(t *testing.T) {
	filename := "a/b:cde"

	dir := filepath.Join("tmp", filename)
	if err := os.MkdirAll(dir, 0777); err == nil {
		t.Fatalf("%v is safe string.", filename)
	}

	safeFilename := SafeFilename(filename)
	safeDir := filepath.Join("tmp", safeFilename)
	if err := os.MkdirAll(safeDir, 0777); err != nil {
		t.Fatalf("%v", err)
	}
}
