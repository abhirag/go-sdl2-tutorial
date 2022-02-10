package main

import (
	"os"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	bmpImagePath = "./assets/stretch.bmp"
)

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface
var gStretchedSurface *sdl.Surface

func initiate() (err error) {
	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}

	if err = img.Init(img.INIT_PNG); err != nil {
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

	if loadedSurface, err = sdl.LoadBMP(path); err != nil {
		return nil, err
	}
	defer loadedSurface.Free()

	if optimizedSurface, err = loadedSurface.Convert(gScreenSurface.Format, 0); err != nil {
		return nil, err
	}

	return optimizedSurface, nil
}

func loadMedia() (err error) {
	gStretchedSurface, err = loadSurface(bmpImagePath)
	return err
}

func terminate() (err error) {
	gStretchedSurface.Free()

	if err = gWindow.Destroy(); err != nil {
		return err
	}

	img.Quit()
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
		var stretchRect sdl.Rect
		stretchRect.X = 0
		stretchRect.Y = 0
		stretchRect.W = screenWidth
		stretchRect.H = screenHeight
		if err = gStretchedSurface.BlitScaled(nil, gScreenSurface, &stretchRect); err != nil {
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
