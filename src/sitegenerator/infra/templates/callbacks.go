package templates

type FuncCallbacks interface {
	ContentRoot() string
	AddAssetHash(urlPath string, hash string) string
}

func CreateFuncCallbacks(contentRoot string) FuncCallbacks {
	return &defaultFuncCallbacks{
		contentRoot: contentRoot,
	}
}

type defaultFuncCallbacks struct {
	contentRoot string
}

func (fo *defaultFuncCallbacks) ContentRoot() string {
	return fo.contentRoot
}

func (fo *defaultFuncCallbacks) AddAssetHash(urlPath string, hash string) string {
	return urlPath + "?v=" + hash
}
