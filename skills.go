package main

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

func skills() app.HTMLDiv {
	return app.Div().Class("skills-wrapper").Body(
		app.Div().Class("first-set").Body(
			app.Br(),
			app.H3().Text("Mobile"),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/java-original.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/kotlin-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/android-original.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/gradle-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/git-original.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/firebase-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
		),
		app.Div().Class("second-set").Body(
			app.Br(),
			app.H3().Text("Backend"),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/go-original-wordmark.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/postgresql-original.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/redis-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/linux-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/docker-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/circleci-plain.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
		),
		app.Div().Class("third-set").Body(
			app.Br(),
			app.H3().Text("Natural Languages"),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/uz.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/gb.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
				),
			),
			app.Div().Class("icon-card").Body(
				app.Img().Class("icon").Src("/web/images/ru.svg").Alt(""),
				app.Div().Class("progress-segment").Body(
					app.Div().Class("item active-segment"),
					app.Div().Class("item active-segment"),
					app.Div().Class("item"),
					app.Div().Class("item"),
				),
			),
		),
	)
}
