package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"time"
)

type Project struct {
	Title     string
	Image     string
	StartDate time.Time
	EndDate   time.Time
	IsActive  bool
	About     []string
	Link      app.UI
}

var projects = []Project{
	{
		Title:     "Comfort-Oil User",
		Image:     "web/images/comfort_user.png",
		StartDate: yearMonthTime(2021, time.September),
		EndDate:   yearMonthTime(2021, time.November),
		IsActive:  true,
		About: []string{
			"Android app for users of the company which ships oil throughout Uzbekistan.",
			"Used technologies: Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager",
		},
		Link: app.Div().Body(
			app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
			app.Div().Style("display", "inline").Style("width", "8px"),
			app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.comfortoil_user"),
		),
	},
	{
		Title:     "Comfort-Oil Driver",
		Image:     "web/images/comfort_driver.png",
		StartDate: yearMonthTime(2021, time.September),
		EndDate:   yearMonthTime(2021, time.November),
		IsActive:  true,
		About: []string{
			"Android app for the drivers of the company which ships oil throughout Uzbekistan.",
			"Used technologies: Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager",
		},
		Link: app.Div().Body(
			app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
			app.Div().Style("display", "inline").Style("width", "8px"),
			app.A().Style("display", "inline").Class("text").Text("Application distributed among drivers only (no link 😑)").Href("#"),
		),
	},
	{
		Title:     "Edu Tizim Teacher",
		Image:     "web/images/edu_system_teacher.png",
		StartDate: yearMonthTime(2021, time.December),
		EndDate:   yearMonthTime(2022, time.February),
		IsActive:  true,
		About: []string{
			"Android app for managing Learning Centers and Schools of any kind.",
			"Used technologies: Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager",
		},
		Link: app.Div().Body(
			app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
			app.Div().Style("display", "inline").Style("width", "8px"),
			app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.edusystem.teacher"),
		),
	},
	{
		Title:     "Edu Tizim Parent",
		Image:     "web/images/edu_system_parent.png",
		StartDate: yearMonthTime(2022, time.February),
		EndDate:   yearMonthTime(2022, time.April),
		IsActive:  true,
		About: []string{
			"Android app for parents to monitor their children' studies.",
			"Used technologies: Datastore, Paging-3, Room, Navigation-component, Hilt, Retrofit, Kotlin-flow",
			"Jetpack compose",
		},
		Link: app.Div().Body(
			app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
			app.Div().Style("display", "inline").Style("width", "8px"),
			app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.edusystem.parent"),
		),
	},
}

type portfolioPage struct {
	app.Compo
	projects []Project
}

func newPortfolioPage() *portfolioPage {
	return &portfolioPage{}
}

func (p *portfolioPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *portfolioPage) OnNav(ctx app.Context) {
	p.projects = projects
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
		Icon(logoSVG).
		Content(
			generateProjects(p.projects)...,
		)
}

func generateProjects(projects []Project) []app.UI {
	var data []app.UI
	for _, project := range projects {
		supporting := ""
		if project.IsActive {
			supporting = " ∘ Still supporting"
		}
		data = append(data, app.Section().Body(
			app.H2().Text(project.Title),
			app.P().Class("text").Text(
				fmt.Sprintf("%s - %s%s",
					project.StartDate.Format("Jan 2006"),
					project.EndDate.Format("Jan 2006"),
					supporting,
				),
			),
			app.Br(),
			app.Img().Class("project-banner").Src(project.Image),
			app.Br(),
			app.Ul().Body(
				app.Range(project.About).Slice(func(i int) app.UI {
					return app.Li().Body(
						app.P().Class("text").Text(project.About[i]).Style("font-size", "16px"),
					)
				}),
			),
			app.Br(),
			project.Link,
		),
			app.Br(),
			app.Br(),
		)
	}
	return data
}
