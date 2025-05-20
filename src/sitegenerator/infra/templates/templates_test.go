package templates

import (
	"testing"

	"sitegenerator/app"
	"sitegenerator/infra/testdata"

	"github.com/stretchr/testify/assert"
)

func TestRenderArticle(t *testing.T) {
	templates, err := parseSiteTemplates()
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := templates.GenerateArticlePage(&app.ArticlePageDetails{
		IsVisible: true,
		Meta:      testdata.ExpectedMetadata("internal/markdown-demo.html"),
		Content:   []byte("<strong>Hello, World!</strong>"),
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := testdata.ExpectedHtml("internal/markdown-demo.html")
	assert.Equal(t, expected, string(bytes))
}

func TestRenderSection(t *testing.T) {
	templates, err := parseSiteTemplates()
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := templates.GenerateSectionPage(&app.SectionPageDetails{
		Title:     "Golang",
		IsVisible: true,
		Pages: []app.UrlAndValue[*app.ArticleMetadata]{
			{
				Url:   "/golang/error-handling.html",
				Value: testdata.ExpectedMetadata("golang/error-handling.html"),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := testdata.ExpectedHtml("golang.html")
	assert.Equal(t, expected, string(bytes))
}

// TODO: Тест рендеринга главной страницы (index)
func TestRenderIndex(t *testing.T) {
	templates, err := parseSiteTemplates()
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := templates.GenerateIndexPage(&app.IndexPageData{
		Sections: []app.UrlAndValue[string]{
			{
				Url:   "/golang.html",
				Value: "Golang",
			},
			{
				Url:   "/atdd.html",
				Value: "ATDD",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := testdata.ExpectedHtml("index.html")
	assert.Equal(t, expected, string(bytes))
}

func parseSiteTemplates() (app.SiteTemplates, error) {
	callbacks := CreateFuncCallbacks(testdata.PublicDir())
	return ParseSiteTemplates(callbacks, testdata.TemplatesDir())
}
