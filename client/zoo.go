package client

import (
	"encoding/json"
	"fmt"
	"image/color"
	_ "image/png" //some comment for the linter
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	gamestate "github.com/chafgames/chaos-monkey/gamestate"
)

var (
	win     *pixelgl.Window
	binPath string

	playerPics  []*pixel.Sprite
	playerSize  = pixel.V(82, 100)
	playerSpeed = 100.0

	camSpeed     = 500.0
	camZoom      = 1.0
	camZoomSpeed = 1.2
)

func run() {
	var err error
	fmt.Println("Started...")
	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 512, 360),
		VSync:  true,
	}

	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	camPos := win.Bounds().Center()
	playerVec := win.Bounds().Center()

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
		updateState()
		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.Pressed(pixelgl.KeyLeft) {
			playerVec.X -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			playerVec.X += playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			playerVec.Y -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			playerVec.Y += playerSpeed * dt
		}

		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		// fmt.Println(playerPos)
		win.Clear(color.Black)
		// Draw all layers to the window.
		// matLevel := pixel.IM
		// mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
		// matLevel = matLevel.ScaledXY(pixel.ZV, pixel.V(5, 5))
		// matLevel = matLevel.ScaledXY(pixel.ZV, pixel.V(2, 2))
		// matLevel = matLevel.Moved(pixel.ZV)

		if err := m.DrawAll(win, color.Black, pixel.IM.Moved(pixel.ZV)); err != nil {
			panic(err)
		}

		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))
		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))

		// mat := pixel.IM
		//update local IM state for drawing

		_, ok := state.Players[myPlayerID]
		if ok {
			state.Players[myPlayerID].IdentityMatrix = state.Players[myPlayerID].IdentityMatrix.Moved(playerVec)
		}

		//TODO: send update of player's IM
		for _, player := range players {
			player.draw()
		}
		// mat = mat.Moved(playerPos)
		// mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
		// mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(5, 5))

		// playerPics[0].Draw(win, mat)

		win.Update()
	}
}

var state *gamestate.GameState
var client *socClient
var myPlayerID string
var players = make(map[string]player)

func updateState() {
	// state.ID++
	client.Notice("update me") // send a msg to trigger broadcast (optional?)
	numAlivePlayers := 0

	for stateID, pState := range state.Players {
		_, foundLocally := players[stateID]
		if foundLocally == false {
			newPlayer := player{
				ID:      stateID,
				Sprites: playerPics,
				Score:   0,
				Health:  100,
			}
			players[newPlayer.ID] = newPlayer
		}
		if pState.Active {
			numAlivePlayers++
			log.Printf("%s:MAT:%v", stateID, pState.IdentityMatrix)

		}
	}
}

//Init - //TODO
func Init() {
	var err error
	state = gamestate.NewGameState()

	myPlayerID = fmt.Sprintf("player-%d", time.Now().Unix()) //TODO: this is where player name would go in?
	client, err = newSocClient(myPlayerID)                   //TODO: err handling
	if err != nil {
		log.Fatalf("Err from server %s", err)
		os.Exit(1)
	}
	client.SocketioClient.On(myPlayerID, func(msg string) {
		// log.Printf("Update from server :%+v\n", msg)
		log.Printf("%s: %s", myPlayerID, msg)
	})
	client.SocketioClient.On("update", func(msg string) {
		// log.Printf("Update from server :%+v\n", msg)
		decodeErr := json.Unmarshal([]byte(msg), &state)
		if decodeErr != nil {
			log.Fatalf("Failed to unmarshal update %+v", state)
			log.Fatalf("jsonErr: %s", decodeErr)
		}
		log.Printf("got update")
	})
	client.SocketioClient.On("disconnection", func(msg string) {
		log.Printf("disconnected: %s", msg)
		os.Exit(1)
	})

	log.Printf("registering as %s", myPlayerID)
	err = client.SocketioClient.Emit("register", myPlayerID)
	// err = client.SocketioClient.Emit("register", myPlayerID)
	// _ = err
}

func shutdown() {
	log.Printf("Shutting down")
	client.Bye(myPlayerID)
	log.Printf("Done")
	return
}

//Run - main game entrypoint
func Run() {
	Init()
	pixelgl.Run(run)

	shutdown()
	log.Println("Done")

}
