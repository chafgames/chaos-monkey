package zoo

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png" //some comment for the linter
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"

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

		imd.Draw(canvas)

		canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Center()))
		win.Update()

	}
}

var state *zoogamestate.GameState
var client *socketio.Client
var myPlayerID string

func updateState() {
	// state.ID++
	client.Notice("update me")
	numAlivePlayers := 0
	for _, player := range state.Players {
		if player.Active {
			numAlivePlayers++
		}
	}
	log.Printf("ACTIVE %d STATE: %+v", numAlivePlayers, state)
}

type someplayerGraphicsTmxStruct struct {
	// this is a dummy which dog will replace when tmx is loaded
}

//Init - //TODO
func Init() {
	var err error
	state = zoogamestate.NewGameState()

	myPlayerID = fmt.Sprintf("player-%d", time.Now().Unix()) //TODO: this is where player name would go in?
	client, err = socketio.NewClient(myPlayerID)             //TODO: err handling
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
