package cueutil

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CueSuffix  = ".cue"
	ModDir     = "cue.mod"
	ConfigFile = "module.cue"
	PkgDir     = "pkg"
)

var ErrFileNotFound = errors.New("file not found")

// FindModDirPath finds `cue.mod` directory. empty string is returned if not found.
func FindModDirPath() (path string) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	stop := "/"
	dir, err := findParentFile(ModDir, pwd, stop)
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ModDir)
}

func findParentFile(file, start, stop string) (string, error) {
	files, err := ioutil.ReadDir(start)
	if err != nil {
		return "", err
	}

	if dirHasFile(files, file) {
		return start, nil
	} else if start == stop {
		return "", ErrFileNotFound
	}

	return findParentFile(file, filepath.Dir(start), stop)
}

func dirHasFile(files []os.FileInfo, filename string) bool {
	for _, f := range files {
		if f.Name() == filename {
			return true
		}
	}

	return false
}
