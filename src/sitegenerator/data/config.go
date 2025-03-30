package data

import (
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	SourceDir    string
	TargetDir    string
	TemplatesDir string
}

type ConfigFileData struct {
	SourceDir    string `yaml:"sourceDir"`
	TargetDir    string `yaml:"targetDir"`
	TemplatesDir string `yaml:"templatesDir"`
}

func ReadConfig(path string) (*Config, error) {
	baseDir := filepath.Dir(path)
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configData := ConfigFileData{
		SourceDir:    "content",
		TargetDir:    "public",
		TemplatesDir: "templates",
	}

	err = yaml.Unmarshal(contents, &configData)
	if err != nil {
		return nil, err
	}

	return &Config{
		SourceDir:    filepath.Join(baseDir, configData.SourceDir),
		TargetDir:    filepath.Join(baseDir, configData.TargetDir),
		TemplatesDir: filepath.Join(baseDir, configData.TemplatesDir),
	}, nil
}
