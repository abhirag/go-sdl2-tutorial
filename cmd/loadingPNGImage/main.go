package main

import (
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	pngImagePath = "./assets/loaded.png"
)

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface
var gPNGSurface *sdl.Surface

func initiate() (err error) {
	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}

	if gWindow, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN); err != nil {
		return err
	}

	if gScreenSurface, err = gWindow.GetSurface(); err != nil {
		return err
	}

	return nil
}

func loadSurface(path string) (optimizedSurface *sdl.Surface, err error) {
	var loadedSurface *sdl.Surface

	if loadedSurface, err = img.Load(path); err != nil {
		return nil, err
	}
	defer loadedSurface.Free()

	if optimizedSurface, err = loadedSurface.Convert(gScreenSurface.Format, 0); err != nil {
		return nil, err
	}

	return optimizedSurface, nil
}

func loadMedia() (err error) {
	gPNGSurface, err = loadSurface(pngImagePath)
	return err
}

func terminate() (err error) {
	gPNGSurface.Free()

	if err = gWindow.Destroy(); err != nil {
		return err
	}

	sdl.Quit()

	return nil
}

func run() (err error) {
	if err = initiate(); err != nil {
		return err
	}

	if err = loadMedia(); err != nil {
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
		if err = gPNGSurface.Blit(nil, gScreenSurface, nil); err != nil {
			return err
		}

		if err = gWindow.UpdateSurface(); err != nil {
			return err
		}
	}

	err = terminate()
	return err
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
