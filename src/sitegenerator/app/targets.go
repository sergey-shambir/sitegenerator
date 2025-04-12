package app

type Targets interface {
	Write(path string, data []byte) error
}
