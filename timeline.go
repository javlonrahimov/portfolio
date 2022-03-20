package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func timeline() app.HTMLDiv {
	return app.Div().Class("timeline").Body(
		app.Br(),
		app.Br(),
		app.Div().Class("container left").Body(
			app.A().
				Style("font-size", "1.2em").
				Text("Unical solutions").
				Href("https://www.linkedin.com/company/unical-solutions/"),
			app.Ul().Body(
				app.Li().Body(
					app.Div().Class("date").Text("Aug 2021 - Present ∘ 8 month"),
				),
				app.Li().Body(
					app.P().Class("text").Text("Android developer"),
				),
				app.Li().Body(
					app.P().Class("text").Text("Right now I am working here with an amazing team, and we have been building challenging and interesting software."),
				),
			),
		),
		app.Br(),
		app.Br(),
		app.Div().Class("container right").Body(
			app.A().
				Style("font-size", "1.2em").
				Text("Invan Systems").
				Href("https://www.linkedin.com/company/invan-systems/"),
			app.Ul().Body(
				app.Li().Body(
					app.Div().Class("date").Text("May 2021 - Aug 2021 ∘ 3 month"),
				),
				app.Li().Body(
					app.P().Class("text").Text("Android developer"),
				),
				app.Li().Body(
					app.P().Class("text").Text("This was my first ever job where I have improved my programming skills and developed a few useful android apps."),
				),
			),
		),
		app.Br(),
		app.Br(),
		app.Div().Class("container right").Body(
			app.A().
				Style("font-size", "1.2em").
				Text("TUIT").
				Href("https://tuit.uz"),
			app.Ul().Body(
				app.Li().Body(
					app.Div().Class("date").Text("Sep 2018 - Present ∘ 4 years"),
				),
				app.Li().Body(
					app.P().Class("text").Text("Student 😉"),
				),
				app.Li().Body(
					app.P().Class("text").Text("I am a final year student at Tashkent University of Information Technologies. I have been learning a lot of interesting and useful things here."),
				),
			),
		),
	)
}
