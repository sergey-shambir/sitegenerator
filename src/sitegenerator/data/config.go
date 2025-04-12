package data

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	SourceDir            string
	TargetDir            string
	TemplatesDir         string
	IgnoreFileExtensions []string
}

type ConfigFileData struct {
	SourceDir            string   `yaml:"sourceDir"`
	TargetDir            string   `yaml:"targetDir"`
	TemplatesDir         string   `yaml:"templatesDir"`
	IgnoreFileExtensions []string `yaml:"ignoreFileExtensions"`
}

func (d *ConfigFileData) validate() error {
	paths := map[string]string{
		"sourceDir":    d.SourceDir,
		"targetDir":    d.TargetDir,
		"templatesDir": d.TemplatesDir,
	}
	var errs []error
	for option, path := range paths {
		if filepath.IsAbs(path) {
			err := xerrors.Errorf("absolute path '%s' not allowed for '%s' option", path, option)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func ReadConfig(path string) (*Config, error) {
	baseDir := filepath.Dir(path)
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, xerrors.Errorf("failed to read config file '%s': %w", path, err)
	}

	configData := ConfigFileData{
		SourceDir:    "content",
		TargetDir:    "public",
		TemplatesDir: "templates",
	}

	err = yaml.Unmarshal(contents, &configData)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse config file '%s': %w", path, err)
	}

	err = configData.validate()
	if err != nil {
		return nil, xerrors.Errorf("invalid config file '%s': %w", path, err)
	}

	return &Config{
		SourceDir:            filepath.Join(baseDir, configData.SourceDir),
		TargetDir:            filepath.Join(baseDir, configData.TargetDir),
		TemplatesDir:         filepath.Join(baseDir, configData.TemplatesDir),
		IgnoreFileExtensions: configData.IgnoreFileExtensions,
	}, nil
}
