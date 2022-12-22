package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"time"
)

type homePage struct {
	app.Compo
	Started time.Time
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.Started = yearMonthTime(2019, time.September)
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	yearDiff, _ := monthYearDiff(p.Started, time.Now())
	return newPage().
		Title("Hey, I'm Javlon").
		Icon(logoSVG).
		Content(
			app.Section().Class("about-me").Body(
				app.H2().Text("About me"),
				app.P().Class("text").Text("Hi, I'm Javlon, and I'm a passionate Android mobile and Golang backend developer.").Style("font-size", "16px"),
				app.P().Class("text").Text(fmt.Sprintf("I have been programming for %d years now, and I have completed many interesting and challenging projects, from commercial apps that is used by thousands, to pet projects that inspire people to learn astronomy.", yearDiff)).Style("font-size", "16px"),
				app.P().Class("text").Text("Apart from programming I enjoy playing ping-pong and chess as a hobby.").Style("font-size", "16px"),
			),
		)
}
