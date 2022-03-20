//todo

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/logs"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	defaultTitle       = "Esoteric"
	defaultDescription = "Esoteric"
	backgroundColor    = "#2e343a"
	goAppGithubURL     = "https://github.com/maxence-charriere/go-app"
	email              = "javlonrahimov1212@gmail.com"
	linkedinUrl        = "https://www.linkedin.com/in/javlon-rahimov-4b06b6207/"
	twitterUrl         = "https://twitter.com/Javlon1212"
	stackoverflowUrl   = "https://stackoverflow.com/users/12153321/javlon"
	githubUrl          = "https://github.com/javlonrahimov"
	leetodeUrl         = "https://leetcode.com/javlon1212/"
)

type localOptions struct {
	Port int `cli:"p" env:"PORTFOLIO_PORT" help:"The port used by the server that serves the PWA."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	ui.BaseHPadding = 42
	ui.BlockPadding = 18
	analytics.Add(analytics.NewGoogleAnalytics())

	app.Route("/", newHomePage())
	app.Route("/my_portfolio", newPortfolioPage())
	app.Route("/blog", newBlogPage())
	app.Route("/contact", newContactPage())

	app.Handle(installApp, handleAppInstall)
	app.Handle(updateApp, handleAppUpdate)
	app.Handle(getMarkdown, handleGetMarkdown)

	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	localOpts := localOptions{Port: 1234}
	cli.Register("local").
		Help(`Launches a server that serves the portfolio in a local environment.`).
		Options(&localOpts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run the portfolio app on GitHub Pages.`).
		Options(&githubOpts)

	h := app.Handler{
		Name:        "Portfolio",
		Title:       defaultTitle,
		Description: defaultDescription,
		Author:      "Javlon Rahimov",
		Image:       "https://go-app.dev/web/images/go-app.pn",
		Keywords: []string{
			"go-app",
			"go",
			"golang",
			"app",
			"pwa",
		},
		BackgroundColor: backgroundColor,
		ThemeColor:      backgroundColor,
		LoadingLabel:    "Wait a minute",
		Scripts: []string{
			"/web/js/prism.js",
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/css/prism.css",
			"/web/css/docs.css",
			"web/css/portfolio.css",
		},
		RawHeaders: []string{
			`<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-1013306768105236" crossorigin="anonymous"></script>`,
			analytics.GoogleAnalyticsHeader("G-SW4FQEM9VM"),
		},
		Resources: app.GitHubPages("portfolio"),
	}

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, localOpts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}
}

func runLocal(ctx context.Context, h *app.Handler, opts localOptions) {
	app.Log(logs.New("starting go-app documentation service").
		Tag("port", opts.Port).
		Tag("version", h.Version),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		err := s.Shutdown(context.Background())
		if err != nil {
			return
		}
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	if err := app.GenerateStaticWebsite(".", h); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Log("command failed:", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
