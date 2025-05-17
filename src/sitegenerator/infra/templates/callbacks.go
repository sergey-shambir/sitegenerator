package templates

type FuncCallbacks interface {
	OutputRoot() string
	AddAssetHash(urlPath string, hash string) string
}

func CreateFuncCallbacks(outputRoot string) FuncCallbacks {
	return &defaultFuncCallbacks{
		outputRoot: outputRoot,
	}
}

type defaultFuncCallbacks struct {
	outputRoot string
}

func (fo *defaultFuncCallbacks) OutputRoot() string {
	return fo.outputRoot
}

func (fo *defaultFuncCallbacks) AddAssetHash(urlPath string, hash string) string {
	return urlPath + "?v=" + hash
}
