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
)

var gWindow *sdl.Window
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
	// noop
	return nil
}

func terminate() (err error) {
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

		fillRect := sdl.Rect{
			X: screenWidth / 4,
			Y: screenHeight / 4,
			W: screenWidth / 2,
			H: screenHeight / 2,
		}
		if err = gRenderer.SetDrawColor(0xFF, 0x00, 0x00, 0xFF); err != nil {
			return err
		}
		if err = gRenderer.FillRect(&fillRect); err != nil {
			return err
		}

		outlineRect := sdl.Rect{
			X: screenWidth / 6,
			Y: screenHeight / 6,
			W: screenWidth * 2 / 3,
			H: screenHeight * 2 / 3,
		}
		if err = gRenderer.SetDrawColor(0x00, 0xFF, 0x00, 0xFF); err != nil {
			return err
		}
		if err = gRenderer.DrawRect(&outlineRect); err != nil {
			return err
		}

		if err = gRenderer.SetDrawColor(0x00, 0x00, 0xFF, 0xFF); err != nil {
			return err
		}
		if err = gRenderer.DrawLine(0, screenHeight/2, screenWidth, screenHeight/2); err != nil {
			return err
		}

		if err = gRenderer.SetDrawColor(0xFF, 0xFF, 0x00, 0xFF); err != nil {
			return err
		}
		for i := 0; i < screenHeight; i += 4 {
			if err = gRenderer.DrawPoint(screenWidth/2, int32(i)); err != nil {
				return err
			}
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
