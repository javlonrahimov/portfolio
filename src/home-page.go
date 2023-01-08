package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"time"
)

type Experience struct {
	Title            string
	JobTitle         string
	StartDate        time.Time
	EndDate          time.Time
	CurrentlyWorking bool
	City             string
	About            []string
}

type Education struct {
	Title             string
	StartDate         time.Time
	EndDate           time.Time
	CurrentlyStudying bool
	City              string
	About             []string
}

type Contact struct {
	title string
	link  string
}

var experiences = []Experience{
	{
		Title:            "Unical",
		JobTitle:         "Android Developer",
		StartDate:        yearMonthTime(2021, time.August),
		EndDate:          time.Now(),
		CurrentlyWorking: true,
		City:             "Tashkent",
		About: []string{
			"Created an Android app for users of the company which ships petrol throughout Uzbekistan.",
			"Developed software for managing Schools and Learning Centers of any kind.",
			"Effectively coded software changes and alterations based on specific design specifications.",
			"Applied the Hexagonal architecture for Android projects.",
		},
	},
	{
		Title:            "Anor Systems",
		JobTitle:         "Android Developer",
		StartDate:        yearMonthTime(2021, time.April),
		EndDate:          yearMonthTime(2021, time.August),
		CurrentlyWorking: false,
		City:             "Tashkent",
		About: []string{
			"Created an app for managing the warehouse.",
			"Implemented the chat feature for the app that is used to manage educational centers.",
			"Refactored old MVC and MVP projects to MVVM.",
		},
	},
}

var educations = []Education{{
	Title:             "Tashkent University of Informational Technologies",
	StartDate:         yearMonthTime(2018, time.September),
	EndDate:           yearMonthTime(2022, time.June),
	CurrentlyStudying: false,
	City:              "Tashkent",
	About: []string{
		"Specification: Software Engineering",
		"GPA: 4.4/5",
	},
}}

var contacts = []Contact{
	{
		title: "Email",
		link:  "mailto: javlonrahimov1212@gmail.com",
	},
	{
		title: "LinkedIn",
		link:  "https://www.linkedin.com/in/javlon-rahimov-4b06b6207/",
	},
	{
		title: "Twitter",
		link:  "https://twitter.com/Javlon1212",
	},
	{
		title: "Stack",
		link:  "https://stackoverflow.com/users/12153321/javlon",
	},
	{
		title: "Github",
		link:  "https://github.com/javlonrahimov",
	},
	{
		title: "Resume",
		link:  "https://resume.io/r/NGTXRtSBM",
	},
}

type homePage struct {
	app.Compo
	started     time.Time
	experiences []Experience
	educations  []Education
	contacts    []Contact
	skills      []string
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.started = yearMonthTime(2019, time.September)
	p.experiences = experiences
	p.educations = educations
	p.contacts = contacts
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	expUI := generateExperience(p.experiences)
	eduUI := generateEducation(p.educations)
	conUI := generateContact(p.contacts)
	yearDiff, _ := monthYearDiff(p.started, time.Now())
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
			app.Div().Class("separator"),
			app.Div().Class("separator"),
			app.H2().Text("Working Experience"),
			app.Range(expUI).Slice(func(i int) app.UI {
				return expUI[i]
			}),
			app.Div().Class("separator"),
			app.Div().Class("separator"),
			app.H2().Text("Education"),
			app.Range(eduUI).Slice(func(i int) app.UI {
				return eduUI[i]
			}),
			app.Br(),
			ui.Flow().
				StretchItems().
				Content(
					app.Range(conUI).Slice(func(i int) app.UI {
						return conUI[i]
					})),
			app.Br(),
		)
}

func generateExperience(experiences []Experience) []app.UI {
	var data []app.UI
	for _, experience := range experiences {
		endDate := experience.EndDate.Format("Jan 2006")
		if experience.CurrentlyWorking {
			endDate = "Present"
		}
		year, month := monthYearDiff(experience.StartDate, experience.EndDate)
		deltaTime := fmt.Sprintf("%d month", month)
		if year > 0 {
			deltaTime = fmt.Sprintf("%d year, %d month", year, month)
		}
		data = append(data, app.Section().Body(
			app.P().Styles(map[string]string{
				"display":         "flex",
				"justify-content": "space-between",
			}).Body(
				app.Span().Body(
					app.P().Text(fmt.Sprintf("%s @ %s", experience.JobTitle, experience.Title)).Style("font-size", "18px"),
					app.P().Class("text").Text(
						fmt.Sprintf("%s - %s - %s",
							experience.StartDate.Format("Jan 2006"),
							endDate,
							deltaTime,
						),
					),
					app.Ul().Body(
						app.Range(experience.About).Slice(func(i int) app.UI {
							return app.Li().Body(
								app.P().Class("text").Text(experience.About[i]).Style("font-size", "16px"),
							)
						}),
					),
				),
				app.Span().Body(
					app.P().Text(experience.City).Style("font-size", "16px"),
				),
			),
			app.Br(),
			app.Hr().Styles(map[string]string{
				"background-color": "rgba(255, 255, 255, 0.08)", "height": "0.5px", "border": "0",
			}),
		),
		)
	}
	return data
}

func generateEducation(educations []Education) []app.UI {
	var data []app.UI
	for _, experience := range educations {
		endDate := experience.EndDate.Format("Jan 2006")
		if experience.CurrentlyStudying {
			endDate = "Present"
		}
		year, month := monthYearDiff(experience.StartDate, experience.EndDate)
		deltaTime := fmt.Sprintf("%d month", month)
		if year > 0 {
			deltaTime = fmt.Sprintf("%d year, %d month", year, month)
		}
		data = append(data, app.Section().Body(
			app.P().Styles(map[string]string{
				"display":         "flex",
				"justify-content": "space-between",
			}).Body(
				app.Span().Body(
					app.P().Text(fmt.Sprintf("%s", experience.Title)).Style("font-size", "18px"),
					app.P().Class("text").Text(
						fmt.Sprintf("%s - %s - %s",
							experience.StartDate.Format("Jan 2006"),
							endDate,
							deltaTime,
						),
					),
					app.Ul().Body(
						app.Range(experience.About).Slice(func(i int) app.UI {
							return app.Li().Body(
								app.P().Class("text").Text(experience.About[i]).Style("font-size", "16px"),
							)
						}),
					),
				),
				app.Span().Body(
					app.P().Text(experience.City).Style("font-size", "16px"),
				),
			),
			app.Br(),
			app.Hr().Styles(map[string]string{
				"background-color": "rgba(255, 255, 255, 0.08)", "height": "0.5px", "border": "0",
			}),
		),
		)
	}
	return data
}

func generateContact(contacts []Contact) []app.UI {

	var data []app.UI
	for _, contact := range contacts {
		data = append(data, app.Span().Body(ui.Link().
			Class("link heading fit unselectable").
			Label(contact.title).
			Href(contact.link)),
		)
	}

	return data
}
