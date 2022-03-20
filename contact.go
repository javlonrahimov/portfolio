package main

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

func contacts() app.HTMLDiv {
	return app.Div().Class("skills-wrapper").Body(
		app.Div().Class("first-set").Body(
			app.Br(),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/at-symbol.svg").Alt(""),
					app.Div().Text("Email").Class("text").Style("text-align", "center"),
				).Href("mailto: "+email),
			),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/linkedin.svg").Alt(""),
					app.Div().Text("LinkedIn").Class("text").Style("text-align", "center"),
				).Href(linkedinUrl),
			),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/twitter.svg").Alt(""),
					app.Div().Text("Twitter").Class("text").Style("text-align", "center"),
				).Href(twitterUrl),
			),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/stack-overflow.svg").Alt(""),
					app.Div().Text("Stack").Class("text").Style("text-align", "center"),
				).Href(stackoverflowUrl),
			),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/git-hub.svg").Alt(""),
					app.Div().Text("GitHub").Class("text").Style("text-align", "center"),
				).Href(githubUrl),
			),
			app.Div().Class("icon-card").Body(
				app.A().Body(
					app.Img().Class("icon").Src("/web/images/leetcode.svg").Alt(""),
					app.Div().Text("Leetcode").Class("text").Style("text-align", "center"),
				).Href(leetodeUrl),
			),
		),
	)
}
