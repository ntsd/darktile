package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/liamg/darktile/internal/app/darktile/config"
	"github.com/liamg/darktile/internal/app/darktile/gui"
	"github.com/liamg/darktile/internal/app/darktile/hinters"
	"github.com/liamg/darktile/internal/app/darktile/termutil"
)

func init() {
	var startupErrors []error
	var fileNotFound *config.ErrorFileNotFound

	conf, err := config.LoadConfig()
	if err != nil {
		if !errors.As(err, &fileNotFound) {
			startupErrors = append(startupErrors, err)
		}
		conf = config.DefaultConfig()
	}

	var theme *termutil.Theme

	theme, err = config.LoadTheme(conf)
	if err != nil {
		if !errors.As(err, &fileNotFound) {
			startupErrors = append(startupErrors, err)
		}
		theme, err = config.DefaultTheme(conf)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to load default theme: %w", err))
		}
	}

	termOpts := []termutil.Option{
		termutil.WithTheme(theme),
	}

	terminal := termutil.New(termOpts...)

	options := []gui.Option{
		gui.WithFontDPI(conf.Font.DPI),
		gui.WithFontSize(conf.Font.Size),
		gui.WithFontFamily(conf.Font.Family),
		gui.WithOpacity(conf.Opacity),
		gui.WithLigatures(conf.Font.Ligatures),
	}

	// load all hinters
	for _, hinter := range hinters.All() {
		options = append(options, gui.WithHinter(hinter))
	}

	g, err := gui.New(terminal, options...)
	if err != nil {
		log.Fatal(err)
	}

	for _, err := range startupErrors {
		g.ShowError(err.Error())
	}
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	mobile.SetGame(g)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
