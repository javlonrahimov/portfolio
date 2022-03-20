package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type portfolioPage struct {
	app.Compo
}

func newPortfolioPage() *portfolioPage {
	return &portfolioPage{}
}

func (p *portfolioPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *portfolioPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *portfolioPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *portfolioPage) Render() app.UI {
	return newPage().
		Title("Portfolio").
		Icon("/web/images/avatar.png").
		Index(
			newIndexLink().Title("Comfort Oil User"),
			app.Br(),
			newIndexLink().Title("Comfort Oil Driver"),
			app.Br(),
			newIndexLink().Title("Edu system Teacher"),
			app.Br(),
			newIndexLink().Title("Edu system Parent"),
			app.Br(),
			newIndexLink().Title("Go app portfolio"),
		).
		Content(
			comfortUser(),
			app.Br(),
			app.Br(),
			comfortDriver(),
			app.Br(),
			app.Br(),
			eduSystemTeacher(),
			app.Br(),
			app.Br(),
			eduSystemParent(),
			app.Br(),
			app.Br(),
			goPortfolio(),
		)
}
