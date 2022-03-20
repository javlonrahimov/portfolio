package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func comfortDriver() app.HTMLSection {
	return app.Section().ID("comfort-oil-driver").Body(
		app.H2().Class("section-title").Text("Comfort Oil - Driver"),
		app.Br(),
		app.Img().Class("project-banner").Src("web/images/comfort_driver.png"),
		app.Br(),
		app.Ul().Body(
			app.Li().Body(
				app.P().Class("text").Text("Sep 2021 - Nov 2021 ∘ Still supporting"),
			),
			app.Li().Body(
				app.P().Class("text").Text("Android app for the drivers of the company which ships oil throughout Uzbekistan."),
			),
			app.Li().Body(
				app.P().Class("text").Text("Used technologies:"),
				app.Li().Body(
					app.P().Class("text").Text("Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager"),
				),
			),
			app.Br(),
			app.Li().Body(
				app.Div().Body(
					app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
					app.Div().Style("display", "inline").Style("width", "8px"),
					app.A().Style("display", "inline").Class("text").Text("Application distributed among drivers only (no link 😑)").Href("#"),
				),
			),
		),
	)
}

func comfortUser() app.HTMLSection {
	return app.Section().ID("comfort-oil-user").Body(
		app.H2().Class("section-title").Text("Comfort Oil - User"),
		app.Br(),
		app.Img().Class("project-banner").Src("web/images/comfort_user.png"),
		app.Br(),
		app.Ul().Body(
			app.Li().Body(
				app.P().Class("text").Text("Sep 2021 - Nov 2021 ∘ Still supporting"),
			),
			app.Li().Body(
				app.P().Class("text").Text("Android app for users of the company which ships oil throughout Uzbekistan."),
			),
			app.Li().Body(
				app.P().Class("text").Text("Used technologies:"),
				app.Li().Body(
					app.P().Class("text").Text("Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager"),
				),
			),
			app.Br(),
			app.Li().Body(
				app.Div().Body(
					app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
					app.Div().Style("display", "inline").Style("width", "8px"),
					app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.comfortoil_user"),
				),
			),
		),
	)
}

func eduSystemTeacher() app.HTMLSection {
	return app.Section().ID("edu-system-teacher").Body(
		app.H2().Class("section-title").Text("Edu system - Teacher"),
		app.Br(),
		app.Img().Class("project-banner").Src("web/images/edu_system_teacher.png"),
		app.Br(),
		app.Ul().Body(
			app.Li().Body(
				app.P().Class("text").Text("Dec 2021 - Feb 2021 ∘ Still supporting"),
			),
			app.Li().Body(
				app.P().Class("text").Text("Android app for managing Learning Centers and Schools of any kind."),
			),
			app.Li().Body(
				app.P().Class("text").Text("Used technologies:"),
				app.Li().Body(
					app.P().Class("text").Text("Datastore, Paging-3, Room, Navigation-component, Databinding, Fragment, Hilt, Socket-io, Retrofit, Kotlin-flow, Work-manager"),
				),
			),
			app.Br(),
			app.Li().Body(
				app.Div().Body(
					app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
					app.Div().Style("display", "inline").Style("width", "8px"),
					app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.edusystem.teacher"),
				),
			),
		),
	)
}

func eduSystemParent() app.HTMLSection {
	return app.Section().ID("edu-system-parent").Body(
		app.H2().Class("section-title").Text("Edu system - Parent"),
		app.Br(),
		app.Img().Class("project-banner").Src("web/images/edu_system_parent.png"),
		app.Br(),
		app.Ul().Body(
			app.Li().Body(
				app.P().Class("text").Text("Feb 2021 ∘ Still supporting"),
			),
			app.Li().Body(
				app.P().Class("text").Text("Android app for parents to monitor their children' studies."),
			),
			app.Li().Body(
				app.P().Class("text").Text("Used technologies:"),
				app.Li().Body(
					app.P().Class("text").Text("Datastore, Paging-3, Room, Navigation-component, Hilt, Retrofit, Kotlin-flow"),
				),
				app.Li().Body(
					app.P().Class("text").Style("color", "deepskyblue").Text("Jetpack Compose"),
				),
			),
			app.Br(),
			app.Li().Body(
				app.Div().Body(
					app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
					app.Div().Style("display", "inline").Style("width", "8px"),
					app.A().Style("display", "inline").Class("text").Text("Play store").Href("https://play.google.com/store/apps/details?id=uz.unical.edusystem.parent"),
				),
			),
		),
	)
}

func goPortfolio() app.HTMLSection {
	return app.Section().ID("go-app-portfolio").Body(
		app.H2().Class("section-title").Text("Go app - Portfolio"),
		app.Br(),
		app.Img().Class("project-banner").Src("web/images/portfolio.png"),
		app.Br(),
		app.Ul().Body(
			app.Li().Body(
				app.P().Class("text").Text("Mart 2021"),
			),
			app.Li().Body(
				app.P().Class("text").Text("Portfolio for myself."),
			),
			app.Li().Body(
				app.P().Class("text").Text("Used technologies:"),
				app.Li().Body(
					app.P().Class("text").Text("Go, Go-app, Css, net/http, markdown"),
				),
			),
			app.Br(),
			app.Li().Body(
				app.Div().Body(
					app.Div().Style("display", "inline").Class("text").Text("Link to project: "),
					app.Div().Style("display", "inline").Style("width", "8px"),
					app.A().Style("display", "inline").Class("text").Text("Github").Href("#"),
				),
			),
		),
	)
}
