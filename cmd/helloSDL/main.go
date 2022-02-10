package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func run() (err error) {
	var window *sdl.Window
	var screenSurface *sdl.Surface

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN); err != nil {
		return err
	}

	if screenSurface, err = window.GetSurface(); err != nil {
		return err
	}

	if err = screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, 0xFF, 0xFF, 0xFF)); err != nil {
		return err
	}

	if err = window.UpdateSurface(); err != nil {
		return err
	}

	quit := false
	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}
	}

	err = window.Destroy()
	return err
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
