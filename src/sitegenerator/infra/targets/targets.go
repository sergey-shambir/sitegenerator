package targets

import (
	"io"
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
	dst, err := t.openOutputFile(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = dst.Write(data)
	if err != nil {
		return xerrors.Errorf("failed to to output file %s: %w", path, err)
	}
	return nil
}

func (t *targets) Copy(path string, src io.Reader) error {
	dst, err := t.openOutputFile(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return xerrors.Errorf("failed to copy to output file %s: %w", path, err)
	}

	return nil
}

func (t *targets) openOutputFile(path string) (io.WriteCloser, error) {
	err := t.ensureParentDirCreated(path)
	if err != nil {
		return nil, err
	}

	absPath, err := t.resolveAbsPath(path)
	if err != nil {
		return nil, err
	}

	dst, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, xerrors.Errorf("failed to open output file for writing %s: %w", path, err)
	}

	return dst, nil
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

	return t.ensureDirCreatedRecursive(absDirPath)
}

func (t *targets) ensureDirCreatedRecursive(dir string) error {
	if _, err := os.Stat(dir); err == nil {
		return nil
	}

	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		return xerrors.Errorf("root directory does not exist %s", dir)
	}

	err := t.ensureDirCreatedRecursive(parentDir)
	if err != nil {
		return err
	}

	err = os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return xerrors.Errorf("failed to create directory %s: %w", dir, err)
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
