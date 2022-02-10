package pathes

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/awisu2/goUtils/files"
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

func TestMakeSafePath(t *testing.T) {
	dir := "tmp"
	file := filepath.Join(dir, "abc.jpg")

	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("%v", err)
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		t.Fatalf("%v", err)
	}

	safeFilename := MakeSafePath(file, MakeSafePathOption{
		IsDirectory:  false,
		FixDuplicate: true,
	})
	if safeFilename != file {
		t.Error(fmt.Errorf("%v, want %v", safeFilename, file))
	}

	files.Save([]byte(""), file)

	safeFilename = MakeSafePath(file, MakeSafePathOption{
		IsDirectory:  false,
		FixDuplicate: true,
	})
	if safeFilename != filepath.Join(dir, "abc_1.jpg") {
		t.Error(fmt.Errorf("%v, want %v", "abc_1.jpg", file))
	}

	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSplitName(t *testing.T) {
	path := "/a/b/c/abc.jpg"
	name, ext := SplitName(filepath.Base(path))
	if name != "abc" {
		t.Error(fmt.Errorf("name: %v, want abc", name))
	}
	if ext != ".jpg" {
		t.Error(fmt.Errorf("ext: %v, want .jpg", ext))
	}
}
