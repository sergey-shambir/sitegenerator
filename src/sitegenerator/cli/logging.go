package cli

import (
	"sitegenerator/app"

	"github.com/sirupsen/logrus"
)

type generatorLogger struct {
	logger *logrus.Logger
}

func newGeneratorLogger() app.GeneratorLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	return &generatorLogger{
		logger: logger,
	}
}

func (g *generatorLogger) LogCopiedFile(path string) {
	g.logger.Info("Copied file " + path)
}

func (g *generatorLogger) LogConvertedFile(path string, outputPath string) {
	g.logger.Info("Converted file " + path + " to " + outputPath)
}

func (g *generatorLogger) LogGeneratedPage(outputPath string) {
	g.logger.Info("Generated page " + outputPath)
}
