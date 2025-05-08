package templates

import (
	"bytes"
	"html/template"
	"path/filepath"
	"sitegenerator/app"
	"strings"

	"golang.org/x/xerrors"
)

type ArticlePageData struct {
	IsVisible bool
	Meta      *app.PageMetadata
	Content   template.HTML
}

type SectionPageItem struct {
	Url  string
	Meta app.PageMetadata
}

type SectionPageData struct {
	IsVisible bool
	Title     string
	Pages     []SectionPageItem
}

type IndexPageItem struct {
	Url   string
	Title string
}

type IndexPageData struct {
	Title    string
	Sections []IndexPageItem
}

type siteTemplates struct {
	tpl *template.Template
}

func ParseSiteTemplates(templatesDir string) (*siteTemplates, error) {

	funcs := template.FuncMap{
		"join":         join,
		"addAssetHash": addAssetHash,
	}

	tpl, err := template.New("sitegenerator").Funcs(funcs).ParseGlob(
		filepath.Join(templatesDir, "*.html"),
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse site templates at '%s': %w", templatesDir, err)
	}

	return &siteTemplates{tpl}, nil
}

func (t *siteTemplates) GenerateArticle() ([]byte, error) {
	return t.generatePage("article_page.html", &ArticlePageData{})
}

func (t *siteTemplates) generatePage(name string, data any) ([]byte, error) {
	var buffer bytes.Buffer
	err := t.tpl.Lookup(name).Execute(&buffer, data)
	if err != nil {
		return nil, xerrors.Errorf("failed to apply site template '%s': %w", name, err)
	}
	return buffer.Bytes(), nil
}

// Соединяет строки, используя указанный разделитель.
func join(texts []string, separator string) string {
	return strings.Join(texts, separator)
}

// Добавляет к URL Path ресурса его hash-сумму,
// чтобы при изменении ресурса кэш браузера не возвращал его старую версию.
func addAssetHash(urlPath string) string {
	return urlPath
}
