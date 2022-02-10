package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth      = 640
	screenHeight     = 480
	defaultImagePath = "./assets/press.bmp"
	upImagePath      = "./assets/up.bmp"
	downImagePath    = "./assets/down.bmp"
	leftImagePath    = "./assets/left.bmp"
	rightImagePath   = "./assets/right.bmp"
)

const (
	keyDefault = iota
	keyUp
	keyDown
	keyLeft
	keyRight
	keyTotal
)

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface
var gCurrentSurface *sdl.Surface
var keyPressSurfaces [keyTotal]*sdl.Surface

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

func loadSurface(path string) (surface *sdl.Surface, err error) {
	if surface, err = sdl.LoadBMP(path); err != nil {
		return nil, err
	}

	return surface, nil
}

func loadMedia() (err error) {
	if keyPressSurfaces[keyDefault], err = loadSurface(defaultImagePath); err != nil {
		return err
	}

	if keyPressSurfaces[keyUp], err = loadSurface(upImagePath); err != nil {
		return err
	}

	if keyPressSurfaces[keyDown], err = loadSurface(downImagePath); err != nil {
		return err
	}

	if keyPressSurfaces[keyLeft], err = loadSurface(leftImagePath); err != nil {
		return err
	}

	if keyPressSurfaces[keyRight], err = loadSurface(rightImagePath); err != nil {
		return err
	}

	return nil
}

func terminate() (err error) {
	for _, surface := range keyPressSurfaces {
		surface.Free()
	}

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
	gCurrentSurface = keyPressSurfaces[keyDefault]
	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				switch keyCode {
				case sdl.K_UP:
					gCurrentSurface = keyPressSurfaces[keyUp]
				case sdl.K_DOWN:
					gCurrentSurface = keyPressSurfaces[keyDown]
				case sdl.K_LEFT:
					gCurrentSurface = keyPressSurfaces[keyLeft]
				case sdl.K_RIGHT:
					gCurrentSurface = keyPressSurfaces[keyRight]
				default:
					gCurrentSurface = keyPressSurfaces[keyDefault]
				}
			}
		}
		if err = gCurrentSurface.Blit(nil, gScreenSurface, nil); err != nil {
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
