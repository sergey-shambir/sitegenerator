package project

import (
	"io/fs"
	"path/filepath"
	"slices"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type sources struct {
	rootPath string
	files    map[app.SourceType][]string
}

func (s *sources) Root() string {
	return s.rootPath
}

func (s *sources) ListFiles(t app.SourceType) []string {
	return s.files[t]
}

func ReadSources(rootPath string, ignoredFileExtensions []string) (app.Sources, error) {
	s := &sources{
		rootPath: rootPath,
		files:    make(map[app.SourceType][]string),
	}

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return xerrors.Errorf("failed to scan '%s': %w", path, err)
		}
		if !d.Type().IsRegular() {
			return nil
		}
		ext := filepath.Ext(path)
		if slices.Contains(ignoredFileExtensions, ext) {
			return nil
		}

		relativePath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return xerrors.Errorf("failed to get relative path from '%s': %w", path, err)
		}

		t := classify(ext)
		s.files[t] = append(s.files[t], relativePath)

		return nil
	})
	return s, err
}

func classify(ext string) app.SourceType {
	switch ext {
	case ".md":
		return app.Markdown
	case ".scss", ".sass":
		return app.Sass
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return app.Image
	case ".css":
		return app.StyleSheet
	case ".js":
		return app.JavaScript
	}
	return app.Unknown
}
