package client

import (
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
	gosocketio "github.com/graarh/golang-socketio"
	socketio "github.com/graarh/golang-socketio"

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

		_, _ = sendUpdateRequest(mySIOClient)
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
var mySIOClient *socketio.Client
var myPlayerID string   // ID used for comms only
var myOnHands *player   // local player for oncall
var myMonkeys []*player // local players for monkeys
var myPlayer *player    //local player for input mapping

// var myPlayers sync.Map

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
	mySIOClient, err = newSIOClient()                        //TODO: err handling
	if err != nil {
		log.Fatalf("Err from server %s", err)
		os.Exit(1)
	}
	sendJoin(mySIOClient)
	mySIOClient.On(myPlayerID, func(msg string) { //register mmy player id in case we want 1:1
		log.Printf("%s: %s", myPlayerID, msg)
	})

	mySIOClient.On("/updatestate", func(h *gosocketio.Channel, args Message) {
		log.Printf("/updatestate: %s.", args.Text)
		log.Printf("/updatestate: %s.", args.Text)
		log.Printf("/updatestate: %s.", args.Text)
		log.Printf("/updatestate: %s.", args.Text)
	})

	var gotRole = false
	for gotRole == false {
		log.Print("requesting player slot")
		serverResp, _ := sendRegister(mySIOClient)
		if serverResp == "\"player\"" {
			gotRole = true
			myPlayer = myOnHands
		} else if strings.HasPrefix(serverResp, "\"monkey") {
			ltrimmed := strings.TrimPrefix(serverResp, "\"monkey")
			rtrimmed := strings.TrimSuffix(ltrimmed, "\"")
			monkeyIdx, strErr := strconv.Atoi(rtrimmed)
			if strErr != nil {
				log.Printf("ERROR: couldn't assign role from server ret:%s", serverResp)
				continue
			}
			gotRole = true
			myPlayer = myMonkeys[monkeyIdx]

		}
		time.Sleep(3 * time.Second)
	}

	log.Print("Today Matthew, I will be " + myPlayer.ID)
	mySIOClient.On("/register", func(msg string) { // ignore /register topic from now on
	})
	mySIOClient.On("/"+myPlayer.ID, func(msg string) { // sub to my role's topic
		log.Printf("/%s: %s", myPlayer.ID, msg)
	})
}

func shutdown() {
	log.Printf("Shutting down")
	mySIOClient.Emit("bye", myPlayerID)
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
