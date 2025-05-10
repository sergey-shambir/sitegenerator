package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	"sitegenerator/app"
	"sitegenerator/infra/config"
	"sitegenerator/infra/convert"
	"sitegenerator/infra/project"
	"sitegenerator/infra/targets"
)

const (
	IndexFileName  = "index.yaml"
	ConfigFileName = "sitegenerator.yaml"
	CacheFileName  = "sitegenerator.cache.json"

	ConverterRootEnvVar = "SITEGENERATOR_CONVERTER_ROOT"
)

func generate(cmd *cobra.Command, args []string) error {
	rootDir, err := os.Getwd()
	if err != nil {
		return xerrors.Errorf("failed to get working directory: %w", err)
	}

	config, err := config.ReadConfig(filepath.Join(rootDir, ConfigFileName))
	if err != nil {
		return err
	}

	proj, err := project.LoadProject(
		config.SourceDir,
		filepath.Join(config.SourceDir, IndexFileName),
		filepath.Join(rootDir, CacheFileName),
	)
	if err != nil {
		return err
	}

	sources, err := project.ReadSources(config.SourceDir, config.IgnoreFileExtensions)
	if err != nil {
		return err
	}

	targets, err := targets.NewTargets(config.TargetDir)
	if err != nil {
		return err
	}

	converterRoot := os.Getenv(ConverterRootEnvVar)
	converter, err := convert.NewConverter(converterRoot)
	if err != nil {
		return err
	}

	logger := newGeneratorLogger()

	generator := app.NewGenerator(sources, targets, converter, proj, logger)

	return generator.Generate()
}
