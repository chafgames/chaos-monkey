package client

import "github.com/faiface/pixel"

import "fmt"

var (
	RedDisks     []*disk
	GreenDisks   []*disk
	BlueDisks    []*disk
	HardDisks    []*disk
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
	for _, c := range RedDisks {
		fmt.Println("Draw disk : ", c)
		redDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range GreenDisks {
		fmt.Println("Draw disk : ", c)
		greenDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range BlueDisks {
		fmt.Println("Draw disk : ", c)
		blueDiskPic.Draw(target, pixel.IM.Moved(c.pos))
	}
	for _, c := range HardDisks {
		fmt.Println("Draw disk : ", c)
		hardDiskPic.Draw(target, pixel.IM.Moved(c.pos))
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
