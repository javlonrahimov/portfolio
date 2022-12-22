package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type contactPage struct {
	app.Compo
}

func newContactPage() contactPage {
	return contactPage{}
}

func (p contactPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p contactPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p contactPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p contactPage) Render() app.UI {
	return newPage().
		Title("Hey, I'm Javlon").
		Icon(logoSVG).
		Content(

			app.Div().Class("separator"),

			newRemoteMarkdownDoc().Src("/web/documents/home.md"),

			app.Div().Class("separator"),
		)
}
