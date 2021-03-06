package client

import (
	"encoding/json"
	"fmt"
	"image/color"
	_ "image/png" //some comment for the linter
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"

	gosocketio "github.com/graarh/golang-socketio"
	socketio "github.com/graarh/golang-socketio"

	gamestate "github.com/chafgames/chaos-monkey/gamestate"
)

var (
	tileMap *tilepix.Map
	win     *pixelgl.Window
	binPath string

	playerPics  []*pixel.Sprite
	playerSize  = pixel.V(64, 64)
	playerSpeed = 400.0
	playerVec   = pixel.ZV

	cam          = pixel.IM
	camZoom      = 1.0
	camZoomSpeed = 1.2
	playerText   = "Welcome to Chaos Monkey!!!!"
	myBasicTxt   *text.Text
)

func resetText(msg string) {
	myLastMsgFrame = myFrameCount
	myBasicTxt.Clear()
	fmt.Fprintf(myBasicTxt, msg)
}
func loadLevel() {
	// Load and initialise the map.
	var err error
	tileMap, err = tilepix.ReadFile("assets/ServerRoom.tmx")
	if err != nil {
		panic(err)
	}

	for _, l := range tileMap.TileLayers {
		l.SetStatic(true)
	}

	if err := tileMap.GenerateTileObjectLayer(); err != nil {
		panic(err)
	}
	for _, og := range tileMap.ObjectGroups {
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
	for _, obj := range tileMap.GetObjectByName("redDisk") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		RedDisks = append(RedDisks, &disk{pos: point})
	}
	for _, obj := range tileMap.GetObjectByName("greenDisk") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		GreenDisks = append(GreenDisks, &disk{pos: point})
	}
	for _, obj := range tileMap.GetObjectByName("blueDisk") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		BlueDisks = append(BlueDisks, &disk{pos: point})
	}
	for _, obj := range tileMap.GetObjectByName("hardDisk") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		HardDisks = append(HardDisks, &disk{pos: point})
	}
	for _, obj := range tileMap.GetObjectByName("redServer") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		RedServers = append(RedServers, &server{pos: point, active: true, ledLoc: obj.Type, onPi: false})
	}
	for _, obj := range tileMap.GetObjectByName("greenServer") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		GreenServers = append(GreenServers, &server{pos: point, active: false, ledLoc: obj.Type, onPi: false})
	}
	for _, obj := range tileMap.GetObjectByName("blueServer") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		BlueServers = append(BlueServers, &server{pos: point, active: false, ledLoc: obj.Type, onPi: false})
	}
	for _, obj := range tileMap.GetObjectByName("hardServer") {
		point, err := obj.GetPoint()
		if err != nil {
			panic(err)
		}
		HardServers = append(HardServers, &server{pos: point, active: false, ledLoc: obj.Type, onPi: false})
	}
}

func run() {
	var err error
	fmt.Println("Started...")
	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	myBasicTxt = text.New(pixel.V(50, 50), basicAtlas)
	resetText(playerText)
	// fmt.Fprintln(basicTxt, "I support multiple lines!")
	// fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	// _, err = http.Post("http://192.168.1.251:5000/text/Chaos-Monkey", "", nil)
	// if err != nil {
	// 	panic(err)
	// }

	playerVec = win.Bounds().Center()
	loadLevel()

	last := time.Now()
	fps := time.Tick(time.Second / 50)

	for !win.Closed() {

		myFrameCount++
		if (myFrameCount - myLastMsgFrame) >= myMsgDuration {
			myBasicTxt.Clear()
		}

		// updateState()
		_, _ = sendUpdateRequest(mySIOClient)
		cb := myPlayer.collisionBox()

		dt := time.Since(last).Seconds()
		last = time.Now()
		// fmt.Println(playerSpeed * dt)

		if win.Pressed(pixelgl.KeyLeft) {
			if myPlayer.State.CurAnim != "W" {
				myPlayer.State.CurAnim = "W" // face west
				myPlayer.LastAnimIdx = 0     // reset anim
			}

			var shouldMove = true
			var futurePos = cb.Moved(pixel.V(-playerSpeed*dt, 0))

			if rectCollides(futurePos) {
				shouldMove = false
			}
			if playerCollides(futurePos) {
				shouldMove = false
				playerVec.X += (playerSpeed * dt) // BUMP RIGHT!
			}
			if shouldMove == true {
				// playerVec.X += (playerSpeed * dt)
				playerVec.X -= playerSpeed * dt
			}
		}
		if win.Pressed(pixelgl.KeyRight) {
			if myPlayer.State.CurAnim != "E" {
				myPlayer.State.CurAnim = "E" // face west
				myPlayer.LastAnimIdx = 0     // reset anim
			}
			var shouldMove = true
			var futurePos = cb.Moved(pixel.V(playerSpeed*dt*2, 0))

			if rectCollides(futurePos) {
				shouldMove = false
			}
			if playerCollides(futurePos) {
				shouldMove = false
				playerVec.X -= (playerSpeed * dt) // BUMP LEFT!
			}
			if shouldMove == true {
				// playerVec.X += (playerSpeed * dt)
				playerVec.X += playerSpeed * dt
			}
		}
		if win.Pressed(pixelgl.KeyDown) {
			if myPlayer.State.CurAnim != "S" {
				myPlayer.State.CurAnim = "S" // face west
				myPlayer.LastAnimIdx = 0     // reset anim
			}
			var shouldMove = true
			var futurePos = cb.Moved(pixel.V(0, -playerSpeed*dt))

			if rectCollides(futurePos) {
				shouldMove = false
			}
			if playerCollides(futurePos) {
				shouldMove = false
				playerVec.Y += (playerSpeed * dt) // BUMP UP!
			}
			if shouldMove == true {
				// playerVec.X += (playerSpeed * dt)
				playerVec.Y -= playerSpeed * dt
			}
		}
		if win.Pressed(pixelgl.KeyUp) {
			if myPlayer.State.CurAnim != "N" {
				myPlayer.State.CurAnim = "N" // face west
				myPlayer.LastAnimIdx = 0     // reset anim
			}
			var shouldMove = true
			var futurePos = cb.Moved(pixel.V(0, playerSpeed*dt))

			if rectCollides(futurePos) {
				shouldMove = false
			}
			if playerCollides(futurePos) {
				shouldMove = false
				playerVec.Y -= (playerSpeed * dt) // BUMP DOWN!
			}
			if shouldMove == true {
				// playerVec.X += (playerSpeed * dt)
				playerVec.Y += playerSpeed * dt
			}
		}

		if win.Pressed(pixelgl.Key0) {
			playerVec = win.Bounds().Center()
			myPlayer.hasDisk = false
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

		if err := tileMap.DrawAll(win, color.Black, pixel.IM.Moved(pixel.ZV)); err != nil {
			panic(err)
		}

		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))
		// pos := cam.Unproject(win.Bounds().Center().Sub(playerPos))

		// mat := pixel.IM
		//update local IM state for drawing

		// log.Printf("updated %s's IM to %+v", myPlayer.ID, myPlayer.State)
		myPlayer.State.IdentityMatrix = pixel.IM.Moved(playerVec)
		cam = pixel.IM.Moved(win.Bounds().Center().Sub(playerVec))
		win.SetMatrix(cam)
		// fmt.Println("PlV:", playerVec)
		// fmt.Println("Cam: ", cam)
		// fmt.Println("Col: ", cb)

		// draw all players
		myOnHands.draw()
		for _, monkey := range myMonkeys {
			monkey.draw()
		}
		if myPlayer.IsMonkey == false {
			myBasicTxt.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(playerVec))
		}

		diskCollision()
		serverCollision()
		drawDisks(win)
		drawServers(win)

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
		<-fps
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
	animMap := loadPlayerSheet()
	myOnHands = &player{ID: "onhands", State: &state.Player, IsMonkey: false, LastAnimIdx: 0, Sprites: animMap}
}

func initMonkeys() {
	monkey0AnimMap, monkey1AnimMap, monkey2AnimMap, monkey3AnimMap, monkey4AnimMap, monkey5AnimMap, monkey6AnimMap, monkey7AnimMap := loadMonkeySheet()
	monkey0 := &player{ID: "monkey0", IsMonkey: true, State: &state.Monkeys[0], LastAnimIdx: 0, Sprites: monkey0AnimMap}
	monkey1 := &player{ID: "monkey1", IsMonkey: true, State: &state.Monkeys[1], LastAnimIdx: 0, Sprites: monkey1AnimMap}
	monkey2 := &player{ID: "monkey2", IsMonkey: true, State: &state.Monkeys[2], LastAnimIdx: 0, Sprites: monkey2AnimMap}
	monkey3 := &player{ID: "monkey3", IsMonkey: true, State: &state.Monkeys[3], LastAnimIdx: 0, Sprites: monkey3AnimMap}
	monkey4 := &player{ID: "monkey4", IsMonkey: true, State: &state.Monkeys[4], LastAnimIdx: 0, Sprites: monkey4AnimMap}
	monkey5 := &player{ID: "monkey5", IsMonkey: true, State: &state.Monkeys[5], LastAnimIdx: 0, Sprites: monkey5AnimMap}
	monkey6 := &player{ID: "monkey6", IsMonkey: true, State: &state.Monkeys[6], LastAnimIdx: 0, Sprites: monkey6AnimMap}
	monkey7 := &player{ID: "monkey7", IsMonkey: true, State: &state.Monkeys[7], LastAnimIdx: 0, Sprites: monkey7AnimMap}
	myMonkeys = []*player{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
}

func initState() {

	var err error
	state = gamestate.NewGameState() // bring on the monkeys!

	// myPlayer.loadPlayerSheet()
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
		jsonErr := json.Unmarshal([]byte(args.Text), &state)
		if jsonErr != nil {
			log.Printf("Error unmarshalling state! %s", jsonErr)
		}
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
	serverResp, _ := sendBye(mySIOClient, myPlayer.ID)
	if serverResp != "\"byenow\"" {
		log.Printf("Server didn't say goodbye :(")
	}
	log.Printf("Done")
	return
}

var mySIOAddr string
var mySIOPort int

//Run - main game entrypoint
func Run(args []string) {
	if len(args) == 0 {
		log.Printf("ERROR: expected at least one arg")
		log.Printf("e.g. ./chaos-monkey 127.0.0.1:3811")
		os.Exit(1)
	}
	if len(args) == 1 && args[0] == "client" {
		log.Printf("ERROR: missing address to connect to")
		log.Printf("e.g. ./chaos-monkey 127.0.0.1:3811")
		os.Exit(1)
	}
	var sioAddrArg string
	if args[0] == "client" {
		sioAddrArg = args[1]
	} else {
		sioAddrArg = args[0]
	}
	var addrReg = regexp.MustCompile(`^[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*$`)
	if addrReg.Match([]byte(sioAddrArg)) == false {
		log.Printf("Error: %s not of expected form e.g. 127.0.0.1:3811", sioAddrArg)
		os.Exit(1)
	}

	splitline := strings.Split(sioAddrArg, ":")
	mySIOAddr = splitline[0]
	mySIOPort, _ = strconv.Atoi(splitline[1])

	initState()
	pixelgl.Run(run)

	shutdown()
	log.Println("Done")

}
