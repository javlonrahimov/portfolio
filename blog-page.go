package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type blogPage struct {
	app.Compo
}

func newBlogPage() *blogPage {
	return &blogPage{}
}

func (p *blogPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *blogPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *blogPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("blog", nil)
}

func (p *blogPage) Render() app.UI {
	return newPage().
		Title("Blog").
		Icon("/web/images/avatar.png").
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
				//newRemoteMarkdownDoc().
				//	Class("fill").
				//	Src("/web/documents/concurrency.md"),
				))
}
