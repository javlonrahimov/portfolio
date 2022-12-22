package main

import (
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// IShell is the interface that describes a layout that responsively displays a
// content with a menu pane, an index page, and a hamburger menu.
type IShell interface {
	app.UI

	// Sets the ID.
	ID(v string) IShell

	// Sets the class. Multiple classes can be defined by successive calls.
	Class(v string) IShell

	// Sets the width in px for the menu and index panes.
	// Default is 270px.
	PaneWidth(px int) IShell

	// Sets the width in px for the ads pane.
	RightPaneWidth(px int) IShell

	// Customizes the hamburger menu button with the given element.
	// Default is ☰.
	HamburgerButton(v app.UI) IShell

	// Sets the hamburger menu content.
	HamburgerMenu(v ...app.UI) IShell

	// Sets the menu pane content.
	Menu(v ...app.UI) IShell

	// Sets the index pane content.
	LeftPane(v ...app.UI) IShell

	// Sets the content.
	Content(v ...app.UI) IShell

	// Sets the ads pane.
	RightPane(v ...app.UI) IShell
}

// Shell returns a layout that responsively displays a content with a menu pane,
// an index pane, and a hamburger menu.
func Shell() IShell {
	return &shell{
		IpaneWidth: 270,
		IadsWidth:  300 + 2*BaseAdHPadding,
		id:         "goapp-shell-" + uuid.NewString(),
	}
}

type shell struct {
	app.Compo

	Iid              string
	Iclass           string
	IpaneWidth       int
	IadsWidth        int
	IhamburgerButton app.UI
	IhamburgerMenu   []app.UI
	Imenu            []app.UI
	IleftPane        []app.UI
	Icontent         []app.UI
	IrightPane       []app.UI

	id                string
	hideMenu          bool
	hideLeftPAne      bool
	hideRightPane     bool
	showHamburgerMenu bool
	width             int
}

func (s *shell) ID(v string) IShell {
	s.Iid = v
	return s
}

func (s *shell) Class(v string) IShell {
	s.Iclass = app.AppendClass(s.Iclass, v)
	return s
}

func (s *shell) PaneWidth(px int) IShell {
	if px > 0 {
		s.IpaneWidth = px
	}
	return s
}

func (s *shell) RightPaneWidth(px int) IShell {
	if px > 0 {
		s.IadsWidth = px
	}
	return s
}

func (s *shell) HamburgerButton(v app.UI) IShell {
	b := app.FilterUIElems(v)
	if len(b) != 0 {
		s.IhamburgerButton = b[0]
	}
	return s
}

func (s *shell) HamburgerMenu(v ...app.UI) IShell {
	s.IhamburgerMenu = app.FilterUIElems(v...)
	return s
}

func (s *shell) Menu(v ...app.UI) IShell {
	s.Imenu = app.FilterUIElems(v...)
	return s
}

func (s *shell) LeftPane(v ...app.UI) IShell {
	s.IleftPane = app.FilterUIElems(v...)
	return s
}

func (s *shell) Content(v ...app.UI) IShell {
	s.Icontent = app.FilterUIElems(v...)
	return s
}

func (s *shell) RightPane(v ...app.UI) IShell {
	s.IrightPane = app.FilterUIElems(v...)
	return s
}

func (s *shell) OnPreRender(ctx app.Context) {
	s.refresh(ctx)
}

func (s *shell) OnMount(ctx app.Context) {
	s.refresh(ctx)
}

func (s *shell) OnResize(ctx app.Context) {
	s.refresh(ctx)
}

func (s *shell) OnUpdate(ctx app.Context) {
	s.refresh(ctx)
}

func (s *shell) Render() app.UI {
	visible := func(v bool) string {
		if v {
			return "block"
		}
		return "none"
	}

	return app.Div().
		DataSet("goapp-ui", "shell").
		ID(s.Iid).
		Class(s.Iclass).
		Body(
			app.Div().
				ID(s.id).
				Style("display", "flex").
				Style("width", "100%").
				Style("height", "100%").
				Style("overflow", "hidden").
				Body(
					app.Div().
						Style("position", "relative").
						Style("display", visible(!s.hideMenu)).
						Style("flex-shrink", "0").
						Style("flex-basis", pxToString(s.IpaneWidth)).
						Style("overflow", "hidden").
						Body(s.Imenu...),
					app.Div().
						Style("position", "relative").
						Style("display", visible(!s.hideLeftPAne)).
						Style("flex-shrink", "0").
						Style("flex-basis", pxToString(s.IpaneWidth)).
						Style("overflow", "hidden").
						Body(s.IleftPane...),
					app.Div().
						Style("position", "relative").
						Style("flex-grow", "1").
						Style("overflow", "hidden").
						Body(s.Icontent...),
					app.Div().
						Style("position", "relative").
						Style("display", visible(!s.hideRightPane)).
						Style("flex-shrink", "0").
						Style("flex-basis", pxToString(s.IadsWidth)).
						Style("overflow", "hidden").
						Body(s.IrightPane...),
				),
			app.Div().
				Style("display", visible(s.hideMenu && len(s.IhamburgerMenu) != 0)).
				Style("position", "absolute").
				Style("top", "0").
				Style("left", "0").
				Style("cursor", "pointer").
				OnClick(s.onHamburgerButtonClick).
				Body(
					app.If(s.IhamburgerButton == nil,
						app.Div().
							Class("goapp-shell-hamburger-button-default").
							Text("☰"),
					),
				),
			app.Div().
				Style("display", visible(s.hideMenu && s.showHamburgerMenu)).
				Style("position", "absolute").
				Style("top", "0").
				Style("left", "0").
				Style("width", "100%").
				Style("height", "100%").
				Style("overflow", "hidden").
				OnClick(s.hideHamburgerMenu).
				Body(s.IhamburgerMenu...),
		)
}

func (s *shell) refresh(ctx app.Context) {
	w, _ := s.layoutSize()

	cw := int(float64(s.IpaneWidth) * 2.70)

	adsExists := len(s.IrightPane) != 0
	hideAds := true
	if adsExists && cw+s.IadsWidth <= w {
		hideAds = false
		cw += s.IadsWidth
	}

	hideIndex := true
	if (!adsExists || !hideAds) && len(s.IleftPane) != 0 && cw+s.IpaneWidth <= w {
		hideIndex = false
		cw += s.IpaneWidth
	}

	hideMenu := true
	if (len(s.IleftPane) == 0 || !hideIndex) && len(s.Imenu) != 0 && cw+s.IpaneWidth <= w {
		hideMenu = false
		cw += s.IpaneWidth
	}

	if hideMenu != s.hideMenu ||
		hideIndex != s.hideLeftPAne ||
		hideAds != s.hideRightPane ||
		w != s.width {
		s.hideMenu = hideMenu
		s.hideLeftPAne = hideIndex
		s.hideRightPane = hideAds
		s.width = w

		ctx.Defer(func(app.Context) {
			s.ResizeContent()
		})
	}
}

func (s *shell) layoutSize() (int, int) {
	layout := app.Window().GetElementByID(s.id)
	if !layout.Truthy() {
		return 320, 568
	}
	return layout.Get("clientWidth").Int(), layout.Get("clientHeight").Int()
}

func (s *shell) onHamburgerButtonClick(ctx app.Context, e app.Event) {
	s.showHamburgerMenu = true
}

func (s *shell) hideHamburgerMenu(ctx app.Context, e app.Event) {
	s.showHamburgerMenu = false
}
