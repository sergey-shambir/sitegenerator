package data

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/xerrors"
	yaml "gopkg.in/yaml.v3"
)

type PageMetadata struct {
	Title       string
	Description string
	Category    string
	Keywords    []string
}

const metadataMarker = "---"

func ParsePageMetadata(path string) (*PageMetadata, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, xerrors.Errorf("failed to open file '%s': %w", path, err)
	}
	contents, err := readPageMetadataYaml(file)
	if err != nil {
		return nil, xerrors.Errorf("failed to read '%s' metadata: %w", path, err)
	}

	var metadata PageMetadata
	err = yaml.Unmarshal([]byte(contents), &metadata)

	return &metadata, err
}

func readPageMetadataYaml(file io.Reader) (string, error) {
	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	isStarted := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == metadataMarker {
			if isStarted {
				return sb.String(), nil
			}
			isStarted = true
		} else {
			sb.WriteString(line)
			sb.WriteRune('\n')
		}
	}

	// NOTE: блок метаданных не был закрыт
	missingHint := "opening marker"
	if isStarted {
		missingHint = "closing marker"
	}

	return "", fmt.Errorf("No YAML %s %s up to file end", missingHint, metadataMarker)
}
