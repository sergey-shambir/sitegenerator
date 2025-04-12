package app

import "fmt"

type Generator struct {
	sources Sources
	targets Targets
	cache   GeneratorCache
}

func NewGenerator(sources Sources, targets Targets, cache GeneratorCache) *Generator {
	return &Generator{
		sources: sources,
		targets: targets,
		cache:   cache,
	}
}

func (g *Generator) Generate() error {
	for i, path := range g.sources.ListFiles(Markdown) {
		fmt.Printf("%d. %s\n", i, path)
	}

	err := g.cache.SaveCache()
	if err != nil {
		return err
	}
	return nil
}
