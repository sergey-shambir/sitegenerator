package app

type GeneratorLogger interface {
	LogCopiedFile(path string)
	LogConvertedFile(path string, outputPath string)
}
