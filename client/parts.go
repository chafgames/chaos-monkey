package client

import "github.com/faiface/pixel"

import "fmt"

var (
	Disks        []*disk
	redDisk, _   = loadPicture("assets/floppy.png")
	greenDisk, _ = loadPicture("assets/floppy_2.png")
	blueDisk, _  = loadPicture("assets/floppy_3.png")
	hardDisk, _  = loadPicture("assets/hdd.png")
	redDiskPic   = pixel.NewSprite(redDisk, redDisk.Bounds())
	greenDiskPic = pixel.NewSprite(greenDisk, greenDisk.Bounds())
	blueDiskPic  = pixel.NewSprite(blueDisk, blueDisk.Bounds())
	hardDiskPic  = pixel.NewSprite(hardDisk, hardDisk.Bounds())
)

type disk struct {
	pos   pixel.Vec
	image pixel.Sprite
}

func drawDisks(target pixel.Target) {
	for _, c := range Disks {
		fmt.Println("Draw disk : ", c)
		redDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
}

func diskCollision() {
	for i, c := range Disks {
		if myPlayer.collisionBox().Contains(c.pos.Add(pixel.V(8, 8))) {
			// Delete disk
			copy(Disks[i:], Disks[i+1:])
			Disks[len(Disks)-1] = nil
			Disks = Disks[:len(Disks)-1]

			// addDisks(5)

			return
		}
	}
}
