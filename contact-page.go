package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type contactPage struct {
	app.Compo
}

func newContactPage() *contactPage {
	return &contactPage{}
}

func (p *contactPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *contactPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *contactPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("contact", nil)
}

func (p *contactPage) Render() app.UI {
	return newPage().
		Title("Contact").
		Icon("/web/images/avatar.png").
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					contacts(),
				))
}
