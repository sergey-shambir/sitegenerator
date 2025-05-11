package templates

import (
	"html/template"

	"sitegenerator/app"
)

type articlePageVars struct {
	IsVisible bool
	Meta      *app.ArticleMetadata
	Content   template.HTML
}

func toArticlePageVars(d app.ArticlePageDetails) *articlePageVars {
	return &articlePageVars{
		IsVisible: d.IsVisible,
		Meta:      d.Meta,
		Content:   template.HTML(d.Content),
	}
}

type sectionPageVars struct {
	IsVisible bool
	Title     string
	Pages     []app.UrlAndValue[*app.ArticleMetadata]
}

func toSectionPageVars(d app.SectionPageDetails) *sectionPageVars {
	return &sectionPageVars{
		IsVisible: d.IsVisible,
		Title:     d.Title,
		Pages:     d.Pages,
	}
}

type indexPageVars struct {
	Title    string
	Sections []app.UrlAndValue[string]
}

func toIndexPageVars(d app.IndexPageData) *indexPageVars {
	return &indexPageVars{
		Title:    d.Title,
		Sections: d.Sections,
	}
}
