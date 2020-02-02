package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/faiface/pixel"
)

var (
	RedDisks       []*disk
	GreenDisks     []*disk
	BlueDisks      []*disk
	HardDisks      []*disk
	RedServers     []*server
	GreenServers   []*server
	BlueServers    []*server
	HardServers    []*server
	redDisk, _     = loadPicture("assets/floppy.png")
	greenDisk, _   = loadPicture("assets/floppy_2.png")
	blueDisk, _    = loadPicture("assets/floppy_3.png")
	hardDisk, _    = loadPicture("assets/hdd.png")
	redDiskPic     = pixel.NewSprite(redDisk, redDisk.Bounds())
	greenDiskPic   = pixel.NewSprite(greenDisk, greenDisk.Bounds())
	blueDiskPic    = pixel.NewSprite(blueDisk, blueDisk.Bounds())
	hardDiskPic    = pixel.NewSprite(hardDisk, hardDisk.Bounds())
	redServerPic   = pixel.NewSprite(redDisk, redDisk.Bounds())
	greenServerPic = pixel.NewSprite(greenDisk, greenDisk.Bounds())
	blueServerPic  = pixel.NewSprite(blueDisk, blueDisk.Bounds())
	hardServerPic  = pixel.NewSprite(hardDisk, hardDisk.Bounds())
	scaling        = 0.0
)

type disk struct {
	pos   pixel.Vec
	image pixel.Sprite
}

type server struct {
	pos    pixel.Vec
	image  pixel.Sprite
	active bool
	ledLoc string
	onPi   bool
}

func drawDisks(target pixel.Target) {
	for _, c := range RedDisks {
		redDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range GreenDisks {
		greenDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range BlueDisks {
		blueDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range HardDisks {
		hardDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
}

func displayMatrix(ledLoc string) {
	var matrix [8][8]int
	for i := range matrix {
		for j := range matrix {
			matrix[i][j] = 0
		}
	}
	x, _ := strconv.Atoi(string(ledLoc[0]))
	y, _ := strconv.Atoi(string(ledLoc[1]))
	matrix[x][y] = 1
	fmt.Println(matrix)
	var s string
	for i, _ := range matrix {
		for _, v := range matrix[i] {
			s = s + strconv.Itoa(v)
		}
	}

	resp, err := http.Post("http://192.168.1.251:5000/image/"+s, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
}

func displayLcd(msg string) {
	resp, err := http.Post("http://192.168.1.251:5000/text/"+msg, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
}

func drawServers(target pixel.Target) {
	if myFrameCount%2 == 0 {
		if scaling > 3 {
			scaling = 0
		} else {
			scaling += 0.15
		}
	}
	for _, c := range RedServers {
		if c.active {
			redServerPic.Draw(target, pixel.IM.Scaled(pixel.ZV, scaling).Moved(c.pos))
			if !c.onPi {
				go displayMatrix(c.ledLoc)
				go displayLcd("Get a RED disk")
				c.onPi = true
			}
		}
	}
	for _, c := range GreenServers {
		if c.active {
			greenServerPic.Draw(target, pixel.IM.Scaled(pixel.ZV, scaling).Moved(c.pos))
			if !c.onPi {
				go displayMatrix(c.ledLoc)
				go displayLcd("Get a GREEN disk")
				c.onPi = true
			}
		}
	}
	for _, c := range BlueServers {
		if c.active {
			blueServerPic.Draw(target, pixel.IM.Scaled(pixel.ZV, scaling).Moved(c.pos))
			if !c.onPi {
				go displayMatrix(c.ledLoc)
				go displayLcd("Get a BLUE disk")
				c.onPi = true
			}
		}
	}
	for _, c := range HardServers {
		if c.active {
			hardServerPic.Draw(target, pixel.IM.Scaled(pixel.ZV, scaling).Moved(c.pos))
			if !c.onPi {
				go displayMatrix(c.ledLoc)
				go displayLcd("Get a HARD disk")
				c.onPi = true
			}
		}
	}
}

func serverCollision() {
	for i, c := range RedServers {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(16, 16))) {
			copy(RedServers[i:], RedServers[i+1:])
			RedServers[len(RedServers)-1] = nil
			RedServers = RedServers[:len(RedServers)-1]

			if myPlayer.redDisk && c.active {
				fmt.Println("Fixed red server")
			}

			// addDisks(5)

			return
		}
	}
}

func diskCollision() {
	for i, c := range RedDisks {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(16, 16))) {
			// Delete disk
			copy(RedDisks[i:], RedDisks[i+1:])
			RedDisks[len(RedDisks)-1] = nil
			RedDisks = RedDisks[:len(RedDisks)-1]
			myPlayer.redDisk = true

			// addDisks(5)

			return
		}
	}
	for i, c := range GreenDisks {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(16, 16))) {
			// Delete disk
			copy(GreenDisks[i:], GreenDisks[i+1:])
			GreenDisks[len(GreenDisks)-1] = nil
			GreenDisks = GreenDisks[:len(GreenDisks)-1]
			myPlayer.greenDisk = true

			// addDisks(5)

			return
		}
	}
	for i, c := range BlueDisks {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(16, 16))) {
			// Delete disk
			copy(BlueDisks[i:], BlueDisks[i+1:])
			BlueDisks[len(BlueDisks)-1] = nil
			BlueDisks = BlueDisks[:len(BlueDisks)-1]
			myPlayer.blueDisk = true

			// addDisks(5)

			return
		}
	}
	for i, c := range HardDisks {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(16, 16))) {
			// Delete disk
			copy(HardDisks[i:], HardDisks[i+1:])
			HardDisks[len(HardDisks)-1] = nil
			HardDisks = HardDisks[:len(HardDisks)-1]
			myPlayer.hardDisk = true

			// addDisks(5)

			return
		}
	}
}
