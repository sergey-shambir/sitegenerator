package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	"sitegenerator/app"
	"sitegenerator/data"
)

const (
	ConfigFileName = "sitegenerator.yaml"
	CacheFileName  = "sitegenerator.cache.json"
)

func generate(cmd *cobra.Command, args []string) error {
	rootDir, err := os.Getwd()
	if err != nil {
		return xerrors.Errorf("failed to get working directory: %w", err)
	}

	config, err := data.ReadConfig(filepath.Join(rootDir, ConfigFileName))
	if err != nil {
		return err
	}

	cache, err := data.LoadGeneratorCache(config.SourceDir, filepath.Join(rootDir, CacheFileName))
	if err != nil {
		return err
	}

	sources, err := data.ReadSources(config.SourceDir, config.IgnoreFileExtensions)
	if err != nil {
		return err
	}

	targets, err := data.NewTargets(config.TargetDir)
	if err != nil {
		return err
	}

	generator := app.NewGenerator(sources, targets, cache)

	return generator.Generate()
}
