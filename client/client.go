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
	"strconv"
	"strings"
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
	playerSize  = pixel.V(48, 48)
	playerSpeed = 400.0

	camZoom      = 1.0
	camZoomSpeed = 1.2
)

func loadPlayerSheet() {
	playerSheet, err := loadPicture(filepath.Join(binPath, "assets/monkey.png"))
	if err != nil {
		panic(err)
	}

	playerPics = []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 0)),
	}
}

func run() {
	var err error
	fmt.Println("Started...")
	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 1024, 1024),
		VSync:  true,
	}

	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	camPos := win.Bounds().Center()
	playerVec := win.Bounds().Center()

	// Load and initialise the map.
	m, err := tilepix.ReadFile("assets/ServerRoom.tmx")
	if err != nil {
		panic(err)
	}

	for _, l := range m.TileLayers {
		l.SetStatic(true)
	}

	if err := m.GenerateTileObjectLayer(); err != nil {
		panic(err)
	}
	for _, og := range m.ObjectGroups {
		// only get collision groups
		if og.Name == "objs" {
			continue
		}

		for _, obj := range og.Objects {
			r, err := obj.GetRect()
			if err != nil {
				panic(err)
			}

			collisionRs = append(collisionRs, r)
		}
	}

	last := time.Now()
	for !win.Closed() {
		updateState()
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.Pressed(pixelgl.KeyLeft) {
			playerVec.X -= playerSpeed * dt
			camPos.X -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			playerVec.X += playerSpeed * dt
			camPos.X += playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			playerVec.Y -= playerSpeed * dt
			camPos.Y -= playerSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			playerVec.Y += playerSpeed * dt
			camPos.Y += playerSpeed * dt
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

		// log.Printf("updated %s's IM to %+v", myPlayer.ID, myPlayer.State)
		myPlayer.State.IdentityMatrix = pixel.IM.Moved(playerVec)
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		// draw all players
		myOnHands.draw()
		for _, monkey := range myMonkeys {
			monkey.draw()
		}

		// myPlayers.Range(func(key interface{}, value interface{}) bool {
		// 	p := value.(player)
		// 	p.draw()
		// 	return true
		// })
		// mat = mat.Moved(playerPos)
		// mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
		// mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(5, 5))

		// playerPics[0].Draw(win, mat)

		win.Update()
	}
}

var state *gamestate.GameState
var client *socClient
var myPlayerID string   // ID used for comms only
var myOnHands *player   // local player for oncall
var myMonkeys []*player // local players for monkeys
var myPlayer *player    //local player for input mapping

// var myPlayers sync.Map

func updateState() {
	client.Notice("update me") // send a msg to trigger broadcast (optional?)
	// state.Players.Range(func(key interface{}, value interface{}) bool {
	// 	remotePlayerState := value.(player)
	// 	_, hasLocal := myPlayers.Load(remotePlayerState.ID)
	// 	if hasLocal == false {
	// 		newPlayer := player{
	// 			ID:      remotePlayerState.ID,
	// 			Sprites: playerPics,
	// 			Score:   0,
	// 			Health:  100,
	// 		}
	// 		myPlayers.Store(remotePlayerState.ID, newPlayer)
	// 	}
	// 	return true
	// })
}

func initPlayer() {
	myOnHands = &player{ID: "onhands", State: &state.Player, IsMonkey: false, MonkeyIndex: -1, Sprites: playerPics}
}

func initMonkeys() {
	monkey0 := &player{ID: "monkey0", IsMonkey: true, State: &state.Monkeys[0], MonkeyIndex: 0, Sprites: playerPics}
	monkey1 := &player{ID: "monkey1", IsMonkey: true, State: &state.Monkeys[1], MonkeyIndex: 1, Sprites: playerPics}
	monkey2 := &player{ID: "monkey2", IsMonkey: true, State: &state.Monkeys[2], MonkeyIndex: 2, Sprites: playerPics}
	monkey3 := &player{ID: "monkey3", IsMonkey: true, State: &state.Monkeys[3], MonkeyIndex: 3, Sprites: playerPics}
	monkey4 := &player{ID: "monkey4", IsMonkey: true, State: &state.Monkeys[4], MonkeyIndex: 4, Sprites: playerPics}
	monkey5 := &player{ID: "monkey5", IsMonkey: true, State: &state.Monkeys[5], MonkeyIndex: 5, Sprites: playerPics}
	monkey6 := &player{ID: "monkey6", IsMonkey: true, State: &state.Monkeys[6], MonkeyIndex: 6, Sprites: playerPics}
	monkey7 := &player{ID: "monkey7", IsMonkey: true, State: &state.Monkeys[7], MonkeyIndex: 7, Sprites: playerPics}
	myMonkeys = []*player{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
}

func initState() {

	var err error
	state = gamestate.NewGameState() // bring on the monkeys!

	loadPlayerSheet()
	initPlayer()
	initMonkeys()

	myPlayerID = fmt.Sprintf("player-%d", time.Now().Unix()) //TODO: this is where player name would go in?
	client, err = newSocClient()                             //TODO: err handling
	if err != nil {
		log.Fatalf("Err from server %s", err)
		os.Exit(1)
	}
	client.SocketioClient.On(myPlayerID, func(msg string) {
		log.Printf("%s: %s", myPlayerID, msg)
	})
	client.SocketioClient.On(myPlayerID+"-register", func(msg string) {
		log.Printf("register: %s", msg)
		if msg == "MONKEY-ENGAGED-SIGNAL" {
			return
		} else if strings.HasPrefix(msg, "PLAYER-REGISTERED") {
			// registeredID := strings.Split(msg, ":")[1]
			myPlayer = myOnHands
			return
		} else if strings.HasPrefix(msg, "MONKEY-REGISTERED") {
			monkeyIdx, err := strconv.Atoi(strings.Split(msg, ":")[1])
			if err != nil {
				log.Printf("ERROR: could not convert %s to int index", msg)
				return
			}
			myPlayer = myMonkeys[monkeyIdx]
			return
		}
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
	initState()
	pixelgl.Run(run)

	shutdown()
	log.Println("Done")

}
