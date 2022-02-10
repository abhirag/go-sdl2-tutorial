package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	pngImagePath = "./assets/viewport.png"
)

var gWindow *sdl.Window
var gTexture *sdl.Texture
var gRenderer *sdl.Renderer

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

func loadTexture(path string) (texture *sdl.Texture, err error) {
	var loadedSurface *sdl.Surface

	if loadedSurface, err = img.Load(path); err != nil {
		return nil, err
	}
	defer loadedSurface.Free()

	if texture, err = gRenderer.CreateTextureFromSurface(loadedSurface); err != nil {
		return nil, err
	}

	return texture, nil
}

func loadMedia() (err error) {
	gTexture, err = loadTexture(pngImagePath)
	return err
}

func terminate() (err error) {
	if err = gTexture.Destroy(); err != nil {
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

		topLeftViewport := sdl.Rect{
			X: 0,
			Y: 0,
			W: screenWidth / 2,
			H: screenHeight / 2,
		}
		if err = gRenderer.SetViewport(&topLeftViewport); err != nil {
			return err
		}
		if err = gRenderer.Copy(gTexture, nil, nil); err != nil {
			return err
		}

		topRightViewport := sdl.Rect{
			X: screenWidth / 2,
			Y: 0,
			W: screenWidth / 2,
			H: screenHeight / 2,
		}
		if err = gRenderer.SetViewport(&topRightViewport); err != nil {
			return err
		}
		if err = gRenderer.Copy(gTexture, nil, nil); err != nil {
			return err
		}

		bottomViewport := sdl.Rect{
			X: 0,
			Y: screenHeight / 2,
			W: screenWidth,
			H: screenHeight / 2,
		}
		if err = gRenderer.SetViewport(&bottomViewport); err != nil {
			return err
		}
		if err = gRenderer.Copy(gTexture, nil, nil); err != nil {
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
