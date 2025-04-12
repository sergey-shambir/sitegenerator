package data

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type targets struct {
	root      string
	knownDirs map[string]bool
}

func NewTargets(root string) (app.Targets, error) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil, xerrors.Errorf("target directory does not exist: %s", root)
	}

	return &targets{
		root:      root,
		knownDirs: make(map[string]bool),
	}, nil
}

func (t *targets) Write(path string, data []byte) error {
	err := t.ensureParentDirCreated(path)
	if err != nil {
		return err
	}

	absPath, err := t.resolveAbsPath(path)
	if err != nil {
		return nil
	}

	err = os.WriteFile(absPath, data, 0644)
	if err != nil {
		return xerrors.Errorf("failed to write %s: %w", path, err)
	}
	return nil
}

func (t *targets) ensureParentDirCreated(path string) error {
	dir := filepath.Dir(path)
	if t.knownDirs[dir] {
		return nil
	}

	err := t.ensureDirCreated(dir)
	if err != nil {
		return err
	}

	t.knownDirs[dir] = true
	return nil
}

func (t *targets) ensureDirCreated(dir string) error {
	absDirPath, err := t.resolveAbsPath(dir)
	if err != nil {
		return err
	}

	err = os.Mkdir(absDirPath, 0755)
	if !os.IsExist(err) {
		return xerrors.Errorf("failed to create directory %s: %w", absDirPath, err)
	}
	return nil
}

func (t *targets) resolveAbsPath(path string) (string, error) {
	absPath := filepath.Clean(filepath.Join(t.root, path))
	if !strings.HasPrefix(absPath, t.root) {
		return "", xerrors.Errorf("path should not escape root directory: %s", path)
	}
	return absPath, nil
}
