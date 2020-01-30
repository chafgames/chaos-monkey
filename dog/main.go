package dog

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"time"

	// We must use blank imports for any image formats in the tileset image sources.
	// You will get an error if a blank import is not made; TilePix does not import
	// specific image formats, that is the responsibility of the calling code.
	_ "image/png"

	"github.com/bcvery1/tilepix"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	binPath string

	playerPics  []*pixel.Sprite
	playerSize  = pixel.V(82, 100)
	playerPos   = pixel.ZV
	playerSpeed = 100.0

	camPos       = pixel.ZV
	camSpeed     = 500.0
	camZoom      = 1.0
	camZoomSpeed = 1.2
)

func spritePos(i, j int) pixel.Rect {
	iF := float64(i)
	jF := float64(j)
	r := pixel.R(
		iF*82,
		jF*100,
		(iF+1)*82,
		(jF+1)*100,
	)
	return r
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	fmt.Println("Started...")
	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 512, 360),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tilemapPic, err := loadPicture(filepath.Join(binPath, "assets/monsters.png"))
	if err != nil {
		panic(err)
	}

	playerPics = []*pixel.Sprite{
		pixel.NewSprite(tilemapPic, spritePos(0, 0)),
	}

	// Load and initialise the map.
	m, err := tilepix.ReadFile("assets/512x360.tmx")
	if err != nil {
		panic(err)
	}

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.Pressed(pixelgl.KeyLeft) {
			playerPos.X -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			playerPos.X += playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			playerPos.Y -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			playerPos.Y += playerSpeed * dt
		}

		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		fmt.Println(playerPos)
		win.Clear(color.Black)
		// Draw all layers to the window.
		// matLevel := pixel.IM
		// mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
		// matLevel = matLevel.ScaledXY(pixel.ZV, pixel.V(5, 5))
		// matLevel = matLevel.ScaledXY(pixel.ZV, pixel.V(2, 2))
		// matLevel = matLevel.Moved(pixel.ZV)

		if err := m.DrawAll(win, color.Black, pixel.IM); err != nil {
			panic(err)
		}

		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))
		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))
		mat := pixel.IM
		mat = mat.Moved(playerPos)
		// mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
		// mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(5, 5))

		playerPics[0].Draw(win, mat)

		win.Update()
	}
}

func Run() {
	pixelgl.Run(run)
}
