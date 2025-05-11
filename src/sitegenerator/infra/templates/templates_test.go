package templates

import (
	"sitegenerator/app"
	"sitegenerator/infra/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderArticle(t *testing.T) {
	templates, err := parseSiteTemplates()
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := templates.GenerateArticlePage(app.ArticlePageDetails{
		IsVisible: true,
		Meta:      testdata.ExpectedMetadata("internal/markdown-demo.md"),
		Content:   []byte("<strong>Hello, World!</strong>"),
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := testdata.ExpectedHtml("internal/markdown-demo.html")
	assert.Equal(t, expected, string(bytes))
}

func parseSiteTemplates() (app.SiteTemplates, error) {
	callbacks := CreateFuncCallbacks(testdata.ContentDir())
	return ParseSiteTemplates(callbacks, testdata.TemplatesDir())
}

// TODO: Тест рендеринга раздела (section)

// TODO: Тест рендеринга главной страницы (index)
