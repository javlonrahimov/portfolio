package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type menu struct {
	app.Compo

	Iclass string

	appInstallable bool
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) Class(v string) *menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnNav(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) OnAppInstallChange(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) Render() app.UI {
	linkClass := "link heading fit unselectable"

	isFocus := func(path string) string {
		if app.Window().URL().Path == path {
			return "focus"
		}
		return ""
	}

	return ui.Scroll().
		Class("menu").
		Class(m.Iclass).
		HeaderHeight(headerHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					app.Header().Body(
						app.A().
							Class("heading").
							Class("app-title").
							Href("/").
							Text("Javlon"),
					),
				),
		).
		Content(
			app.Nav().Body(
				app.Div().Class("separator"),
				ui.Link().
					Class(linkClass).
					Label("Home").
					Href("/").
					Class(isFocus("/")),
				app.Div().Class("separator"),
				ui.Link().
					Class(linkClass).
					Label("Portfolio").
					Href("/my_portfolio").
					Class(isFocus("/my_portfolio")),
				app.Div().Class("separator"),
				ui.Link().
					Class(linkClass).
					Label("Blog").
					Href("/blog").
					Class(isFocus("/blog")),
				app.Div().Class("separator"),
				app.Div().Class("separator"),
				ui.Link().
					Class(linkClass).
					Label("Contact").
					Href("/contact").
					Class(isFocus("/contact")),
			),
		).FooterHeight(headerHeight).
		Footer(
			app.P().Body(
				ui.Stack().
					Class("fill").
					Middle().
					Content(
						app.Text("Build with"),
						app.Div().Style("width", "8px"),
						app.A().
							Href(goAppGithubURL).
							Text("Go-app"),
					).Style("text-align", "center"),
			),
		)
}

func (m *menu) installApp(ctx app.Context, e app.Event) {
	ctx.NewAction(installApp)
}
