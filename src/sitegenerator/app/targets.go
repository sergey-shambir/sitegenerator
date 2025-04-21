package app

import "io"

/**
 * Предоставляет возможность записи генерируемых файлов
 */
type Targets interface {
	Write(path string, data []byte) error

	Copy(path string, src io.Reader) error
}
