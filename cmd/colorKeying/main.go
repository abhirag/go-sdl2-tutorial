package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth         = 640
	screenHeight        = 480
	fooImagePath        = "./assets/foo.png"
	backgroundImagePath = "./assets/background.png"
)

var gWindow *sdl.Window
var gRenderer *sdl.Renderer
var gFooTexture texture
var gBackgroundTexture texture

type texture struct {
	sdlTexture *sdl.Texture
	width      int32
	height     int32
}

func (t *texture) destroy() (err error) {
	if t.sdlTexture != nil {
		if err = t.sdlTexture.Destroy(); err != nil {
			return err
		}
	}
	t.width = 0
	t.height = 0
	return nil
}

func (t *texture) loadFromFile(path string) (err error) {
	if err = t.destroy(); err != nil {
		return err
	}
	var newTexture *sdl.Texture
	var loadedSurface *sdl.Surface

	if loadedSurface, err = img.Load(path); err != nil {
		return err
	}
	defer loadedSurface.Free()

	if err = loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 0xFF, 0xFF)); err != nil {
		return err
	}
	if newTexture, err = gRenderer.CreateTextureFromSurface(loadedSurface); err != nil {
		return err
	}
	t.width = loadedSurface.W
	t.height = loadedSurface.H
	t.sdlTexture = newTexture

	return nil
}

func (t *texture) render(x int32, y int32) (err error) {
	renderQuad := sdl.Rect{
		X: x,
		Y: y,
		W: t.width,
		H: t.height,
	}
	if err = gRenderer.Copy(t.sdlTexture, nil, &renderQuad); err != nil {
		return err
	}
	return nil
}

func initiate() (err error) {
	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}

	if err = img.Init(img.INIT_PNG); err != nil {
		return err
	}

	if !sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1") {
		if _, err = fmt.Fprintln(os.Stderr, "Warning: Linear texture filtering not enabled!"); err != nil {
			return err
		}
	}

	if gWindow, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN); err != nil {
		return err
	}

	if gRenderer, err = sdl.CreateRenderer(gWindow, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return err
	}

	if err = gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF); err != nil {
		return err
	}

	return nil
}

func loadMedia() (err error) {
	if err = gFooTexture.loadFromFile(fooImagePath); err != nil {
		return err
	}
	if err = gBackgroundTexture.loadFromFile(backgroundImagePath); err != nil {
		return err
	}

	return nil
}

func terminate() (err error) {
	if err = gFooTexture.destroy(); err != nil {
		return err
	}
	if err = gBackgroundTexture.destroy(); err != nil {
		return err
	}

	if err = gRenderer.Destroy(); err != nil {
		return err
	}

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
		if err = gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF); err != nil {
			return err
		}
		if err = gRenderer.Clear(); err != nil {
			return err
		}

		if err = gBackgroundTexture.render(0, 0); err != nil {
			return err
		}

		if err = gFooTexture.render(240, 190); err != nil {
			return err
		}

		gRenderer.Present()
	}

	err = terminate()
	return err
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
