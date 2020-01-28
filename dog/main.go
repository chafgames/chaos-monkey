package dog

import (
	"image/color"

	// We must use blank imports for any image formats in the tileset image sources.
	// You will get an error if a blank import is not made; TilePix does not import
	// specific image formats, that is the responsibility of the calling code.
	_ "image/png"

	"github.com/bcvery1/tilepix"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 640, 320),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Load and initialise the map.
	m, err := tilepix.ReadFile("assets/PinkPlayer.tmx")
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(color.White)

		// Draw all layers to the window.
		if err := m.DrawAll(win, color.White, pixel.IM); err != nil {
			panic(err)
		}

		win.Update()
	}
}

func Run() {
	pixelgl.Run(run)
}
