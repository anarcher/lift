package app

import (
	"fmt"
	"lift/pkg/cueutil"
	"path/filepath"

	"os"

	"github.com/anarcher/cue-bundler/pkg/cb"
)

type InitOpts struct {
	Path   string
	Module string
}

func NewInitOpts() InitOpts {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return InitOpts{
		Path:   path,
		Module: "",
	}
}

func Init(opts InitOpts) error {
	if err := initModDir(opts.Path, opts.Module); err != nil {
		return err
	}

	modDir := cueutil.FindModDirPath()
	if err := initCB(modDir); err != nil {
		return err
	}

	return nil
}

func initCB(modDir string) error {
	return cb.Init(modDir)
}

func initModDir(path, module string) error {
	mod := filepath.Join(path, "cue.mod")
	_, err := os.Stat(mod)
	if err == nil {
		return fmt.Errorf("cue.mod directory already exists")
	}

	err = os.Mkdir(mod, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	f, err := os.Create(filepath.Join(mod, "module.cue"))
	if err != nil {
		return err
	}
	defer f.Close()

	// Set module even if it is empty, making it easier for users to fill it in.
	_, err = fmt.Fprintf(f, "module: %q\n", module)

	if err = os.Mkdir(filepath.Join(mod, "usr"), 0755); err != nil {
		return err
	}
	if err = os.Mkdir(filepath.Join(mod, "pkg"), 0755); err != nil {
		return err
	}

	return err
}
