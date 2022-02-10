package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	bmpImagePath = "./assets/hello_world.bmp"
)

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface
var gHelloWorld *sdl.Surface

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

func loadMedia() (err error) {
	gHelloWorld, err = sdl.LoadBMP(bmpImagePath)
	return err
}

func terminate() (err error) {
	gHelloWorld.Free()

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

	if err = gHelloWorld.Blit(nil, gScreenSurface, nil); err != nil {
		return err
	}

	if err = gWindow.UpdateSurface(); err != nil {
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

	err = terminate()
	return err
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
