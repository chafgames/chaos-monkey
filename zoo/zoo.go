package zoo

import (
	"encoding/csv"
	"encoding/json"
	"image"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "image/png" //some comment for the linter

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"golang.org/x/image/colornames"

	socketio "github.com/mattmulhern/game-off-2019-scratch/client"
	zoogamestate "github.com/mattmulhern/game-off-2019-scratch/zoogamestate"
)

func loadAnimationSheet(sheetPath, descPath string, frameWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// total hack, nicely format the error at the end, so I don't have to type it every time
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "error loading animation sheet")
		}
	}()

	// open and load the spritesheet
	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	// create a slice of frames inside the spritesheet
	var frames []pixel.Rect
	for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
		frames = append(frames, pixel.R(
			x,
			0,
			x+frameWidth,
			sheet.Bounds().H(),
		))
	}

	descFile, err := os.Open(descPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])

		anims[name] = frames[start : end+1]
	}

	return sheet, anims, nil
}

func run() {
	rand.Seed(time.Now().UnixNano())

	sheet, anims, err := loadAnimationSheet("assets/sheet.png", "assets/sheet.csv", 12)
	_ = sheet
	_ = anims
	if err != nil {
		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// phys := &gopherPhys{
	// 	gravity:   -512,
	// 	runSpeed:  64,
	// 	jumpSpeed: 192,
	// 	rect:      pixel.R(-6, -7, 6, 7),
	// }

	// anim := &gopherAnim{
	// 	sheet: sheet,
	// 	anims: anims,
	// 	rate:  1.0 / 10,
	// 	dir:   +1,
	// }

	// hardcoded level
	// platforms := []platform{
	// 	{rect: pixel.R(-50, -34, 50, -32)},
	// 	{rect: pixel.R(20, 0, 70, 2)},
	// 	{rect: pixel.R(-100, 10, -50, 12)},
	// 	{rect: pixel.R(120, -22, 140, -20)},
	// 	{rect: pixel.R(120, -72, 140, -70)},
	// 	{rect: pixel.R(120, -122, 140, -120)},
	// 	{rect: pixel.R(-100, -152, 100, -150)},
	// 	{rect: pixel.R(-150, -127, -140, -125)},
	// 	{rect: pixel.R(-180, -97, -170, -95)},
	// 	{rect: pixel.R(-150, -67, -140, -65)},
	// 	{rect: pixel.R(-180, -37, -170, -35)},
	// 	{rect: pixel.R(-150, -7, -140, -5)},
	// }
	// for i := range platforms {
	// 	platforms[i].color = randomNiceColor()
	// }

	// gol := &goal{
	// 	pos:    pixel.V(-75, 40),
	// 	radius: 18,
	// 	step:   1.0 / 7,
	// }

	canvas := pixelgl.NewCanvas(pixel.R(-160/2, -120/2, 160/2, 120/2))
	imd := imdraw.New(sheet)
	imd.Precision = 32

	// camPos := pixel.ZV

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		_ = dt
		updateState()
		last = time.Now()

		// lerp the camera position towards the gopher
		// camPos = pixel.Lerp(camPos, phys.rect.Center(), 1-math.Pow(1.0/128, dt))
		// cam := pixel.IM.Moved(camPos.Scaled(-1))
		// canvas.SetMatrix(cam)

		// // slow motion with tab
		// if win.Pressed(pixelgl.KeyTab) {
		// 	dt /= 8
		// }

		// // restart the level on pressing enter
		// if win.JustPressed(pixelgl.KeyEnter) {
		// 	phys.rect = phys.rect.Moved(phys.rect.Center().Scaled(-1))
		// 	phys.vel = pixel.ZV
		// }

		// control the gopher with keys
		// ctrl := pixel.ZV
		// if win.Pressed(pixelgl.KeyLeft) {
		// 	ctrl.X--
		// }
		// if win.Pressed(pixelgl.KeyRight) {
		// 	ctrl.X++
		// }
		// if win.JustPressed(pixelgl.KeyUp) {
		// 	ctrl.Y = 1
		// }

		// // update the physics and animation
		// phys.update(dt, ctrl, platforms)
		// gol.update(dt)
		// anim.update(dt, phys)

		// draw the scene to the canvas using IMDraw
		canvas.Clear(colornames.Black)
		// imd.Clear()
		// for _, p := range platforms {
		// 	p.draw(imd)
		// }
		// gol.draw(imd)
		// anim.draw(imd, phys)
		// imd.Draw(canvas)

		// stretch the canvas to the window
		// win.Clear(colornames.White)
		// win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		// 	math.Min(
		// 		win.Bounds().W()/canvas.Bounds().W(),
		// 		win.Bounds().H()/canvas.Bounds().H(),
		// 	),
		// ).Moved(win.Bounds().Center()))
		canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Center()))
		win.Update()

	}
}

var state *zoogamestate.GameState
var client *socketio.Client
var myPlayer *player

func updateState() {
	// state.ID++
	client.Notice("update me")
	numAlivePlayers := 0
	for _, player := range state.Players {
		if player.Active {
			numAlivePlayers++
		}
	}
	log.Printf("MY_INTERNAL_STATE: %+v", state)
	log.Printf("ACTIVE_PLAYERS: %d", numAlivePlayers)
}

//Init - //TODO
func Init() {
	state = zoogamestate.NewGameState()

	client, _ = socketio.NewClient() //TODO: err handling

	client.SocketioClient.On("update", func(msg string) {
		// log.Printf("Update from server :%+v\n", msg)
		decodeErr := json.Unmarshal([]byte(msg), &state)
		if decodeErr != nil {
			log.Fatalf("Failed to unmarshal update %+v", state)
			log.Fatalf("jsonErr: %s", decodeErr)
		}
	})

	myPlayerID := "bobthebuilder"
	client.SocketioClient.Emit("register", myPlayerID)

}

func shutdown() {
	log.Printf("Shutting down")
	// client.SocketioClient.Emit()
	client.Bye("quit")
	log.Printf("Done")
	return
}

//Run - main game entrypoint
func Run() {
	Init()
	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	// signal.Notify(sigs)
	// go func() {
	// 	sig := <-sigs
	// 	fmt.Println()

	// 	fmt.Println(sig)
	// 	done <- true
	// }()
	// log.Println("Running Zoo...")
	pixelgl.Run(run)

	// <-done
	// log.Println("Signal detected, tidying up")
	shutdown()
	log.Println("Done")

}
