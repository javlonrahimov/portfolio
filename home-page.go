package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Title("Home").
		Icon("/web/images/avatar.png").
		Content(
			app.Section().Class("about-me").Body(
				app.H2().Class("section-title").Text("About me"),
				app.P().Class("text").Text("\tHi, I'm Javlon, and I'm a passionate Android mobile and Golang backend developer."),
				app.P().Class("text").Text("\tI have been programming for 2 years now, and I have completed many interesting and challenging projects, from commercial apps that is used by thousands, to pet projects that inspire people to learn astronomy."),
				app.P().Class("text").Text("\tApart from programming I enjoy playing ping-pong and chess as a hobby."),
			),
			app.Br(),
			app.Section().Class("skills").Body(
				app.H2().Class("section-title").Text("My skills")),
			app.Br(),
			skills(),
			app.Br(),
			app.Section().Class("timeline").Body(
				app.H2().Class("section-title").Text("Timeline"),
				timeline(),
			),
		)
}
