package client

import (
	"image"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
)

func spritePos(i, j int) pixel.Rect {
	iF := float64(i)
	jF := float64(j)
	r := pixel.R(
		iF*48,
		jF*48,
		(iF+1)*48,
		(jF+1)*48,
	)
	return r
}

func playerSpritePos(j, i int) pixel.Rect {
	iF := float64(i)
	jF := float64(j)
	r := pixel.R(
		iF*64,
		jF*64,
		(iF+1)*64,
		(jF+1)*64,
	)
	return r
}
func loadPlayerSheet() map[string][]*pixel.Sprite {

	playerSheet, err := loadPicture(filepath.Join(binPath, "assets/onHands.png"))
	if err != nil {
		panic(err)
	}
	//Monkey 0
	player0NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, playerSpritePos(12, 0)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 1)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 2)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 3)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 6)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 7)),
		pixel.NewSprite(playerSheet, playerSpritePos(12, 8)),
	}
	player0ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, playerSpritePos(9, 0)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 1)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 2)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 3)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 6)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 7)),
		pixel.NewSprite(playerSheet, playerSpritePos(9, 8)),
	}
	player0WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, playerSpritePos(11, 0)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 1)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 2)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 3)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 6)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 7)),
		pixel.NewSprite(playerSheet, playerSpritePos(11, 8)),
	}
	player0SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, playerSpritePos(10, 0)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 1)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 2)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 3)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 6)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 7)),
		pixel.NewSprite(playerSheet, playerSpritePos(10, 8)),
	}
	player0AnimMap := make(map[string][]*pixel.Sprite)
	player0AnimMap["N"] = player0NSprites
	player0AnimMap["E"] = player0ESprites
	player0AnimMap["W"] = player0WSprites
	player0AnimMap["S"] = player0SSprites

	return player0AnimMap
}

func loadMonkeySheet() (
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite,
	map[string][]*pixel.Sprite) {
	playerSheet, err := loadPicture(filepath.Join(binPath, "assets/monkey.png"))
	if err != nil {
		panic(err)
	}

	//Monkey 0
	monkey0NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 0)),
		pixel.NewSprite(playerSheet, spritePos(0, 1)),
		pixel.NewSprite(playerSheet, spritePos(0, 2)),
	}
	monkey0ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(1, 0)),
		pixel.NewSprite(playerSheet, spritePos(1, 1)),
		pixel.NewSprite(playerSheet, spritePos(1, 2)),
	}
	monkey0WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(2, 0)),
		pixel.NewSprite(playerSheet, spritePos(2, 1)),
		pixel.NewSprite(playerSheet, spritePos(2, 2)),
	}
	monkey0SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(3, 0)),
		pixel.NewSprite(playerSheet, spritePos(3, 1)),
		pixel.NewSprite(playerSheet, spritePos(3, 2)),
	}

	//Monkey 1
	monkey1NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 3)),
		pixel.NewSprite(playerSheet, spritePos(0, 4)),
		pixel.NewSprite(playerSheet, spritePos(0, 5)),
	}
	monkey1ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(1, 3)),
		pixel.NewSprite(playerSheet, spritePos(1, 4)),
		pixel.NewSprite(playerSheet, spritePos(1, 5)),
	}
	monkey1WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(2, 3)),
		pixel.NewSprite(playerSheet, spritePos(2, 4)),
		pixel.NewSprite(playerSheet, spritePos(2, 5)),
	}
	monkey1SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(3, 3)),
		pixel.NewSprite(playerSheet, spritePos(3, 4)),
		pixel.NewSprite(playerSheet, spritePos(3, 5)),
	}

	//Monkey 2
	monkey2NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 6)),
		pixel.NewSprite(playerSheet, spritePos(0, 7)),
		pixel.NewSprite(playerSheet, spritePos(0, 8)),
	}
	monkey2ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(1, 6)),
		pixel.NewSprite(playerSheet, spritePos(1, 7)),
		pixel.NewSprite(playerSheet, spritePos(1, 8)),
	}
	monkey2WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(2, 6)),
		pixel.NewSprite(playerSheet, spritePos(2, 7)),
		pixel.NewSprite(playerSheet, spritePos(2, 8)),
	}
	monkey2SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(3, 6)),
		pixel.NewSprite(playerSheet, spritePos(3, 7)),
		pixel.NewSprite(playerSheet, spritePos(3, 8)),
	}

	//Monkey 3
	monkey3NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 9)),
		pixel.NewSprite(playerSheet, spritePos(0, 10)),
		pixel.NewSprite(playerSheet, spritePos(0, 11)),
	}
	monkey3ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(1, 9)),
		pixel.NewSprite(playerSheet, spritePos(1, 10)),
		pixel.NewSprite(playerSheet, spritePos(1, 11)),
	}
	monkey3WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(2, 9)),
		pixel.NewSprite(playerSheet, spritePos(2, 10)),
		pixel.NewSprite(playerSheet, spritePos(2, 11)),
	}
	monkey3SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(3, 9)),
		pixel.NewSprite(playerSheet, spritePos(3, 10)),
		pixel.NewSprite(playerSheet, spritePos(3, 11)),
	}

	//Monkey 4
	monkey4NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(4, 0)),
		pixel.NewSprite(playerSheet, spritePos(4, 1)),
		pixel.NewSprite(playerSheet, spritePos(4, 2)),
	}
	monkey4ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(5, 0)),
		pixel.NewSprite(playerSheet, spritePos(5, 1)),
		pixel.NewSprite(playerSheet, spritePos(5, 2)),
	}
	monkey4WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(6, 0)),
		pixel.NewSprite(playerSheet, spritePos(6, 1)),
		pixel.NewSprite(playerSheet, spritePos(6, 2)),
	}
	monkey4SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(7, 0)),
		pixel.NewSprite(playerSheet, spritePos(7, 1)),
		pixel.NewSprite(playerSheet, spritePos(7, 2)),
	}

	//Monkey 5
	monkey5NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(4, 3)),
		pixel.NewSprite(playerSheet, spritePos(4, 4)),
		pixel.NewSprite(playerSheet, spritePos(4, 5)),
	}
	monkey5ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(5, 3)),
		pixel.NewSprite(playerSheet, spritePos(5, 4)),
		pixel.NewSprite(playerSheet, spritePos(5, 5)),
	}
	monkey5WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(6, 3)),
		pixel.NewSprite(playerSheet, spritePos(6, 4)),
		pixel.NewSprite(playerSheet, spritePos(6, 5)),
	}
	monkey5SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(7, 3)),
		pixel.NewSprite(playerSheet, spritePos(7, 4)),
		pixel.NewSprite(playerSheet, spritePos(7, 5)),
	}

	//Monkey 6
	monkey6NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(4, 6)),
		pixel.NewSprite(playerSheet, spritePos(4, 7)),
		pixel.NewSprite(playerSheet, spritePos(4, 8)),
	}
	monkey6ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(5, 6)),
		pixel.NewSprite(playerSheet, spritePos(5, 7)),
		pixel.NewSprite(playerSheet, spritePos(5, 8)),
	}
	monkey6WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(6, 6)),
		pixel.NewSprite(playerSheet, spritePos(6, 7)),
		pixel.NewSprite(playerSheet, spritePos(6, 8)),
	}
	monkey6SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(7, 6)),
		pixel.NewSprite(playerSheet, spritePos(7, 7)),
		pixel.NewSprite(playerSheet, spritePos(7, 8)),
	}

	//Monkey 7
	monkey7NSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(4, 9)),
		pixel.NewSprite(playerSheet, spritePos(4, 10)),
		pixel.NewSprite(playerSheet, spritePos(4, 11)),
	}
	monkey7ESprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(5, 9)),
		pixel.NewSprite(playerSheet, spritePos(5, 10)),
		pixel.NewSprite(playerSheet, spritePos(5, 11)),
	}
	monkey7WSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(6, 9)),
		pixel.NewSprite(playerSheet, spritePos(6, 10)),
		pixel.NewSprite(playerSheet, spritePos(6, 11)),
	}
	monkey7SSprites := []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(7, 9)),
		pixel.NewSprite(playerSheet, spritePos(7, 10)),
		pixel.NewSprite(playerSheet, spritePos(7, 11)),
	}

	monkey0AnimMap := make(map[string][]*pixel.Sprite)
	monkey0AnimMap["N"] = monkey0NSprites
	monkey0AnimMap["E"] = monkey0ESprites
	monkey0AnimMap["W"] = monkey0WSprites
	monkey0AnimMap["S"] = monkey0SSprites
	monkey1AnimMap := make(map[string][]*pixel.Sprite)
	monkey1AnimMap["N"] = monkey1NSprites
	monkey1AnimMap["E"] = monkey1ESprites
	monkey1AnimMap["W"] = monkey1WSprites
	monkey1AnimMap["S"] = monkey1SSprites
	monkey2AnimMap := make(map[string][]*pixel.Sprite)
	monkey2AnimMap["N"] = monkey2NSprites
	monkey2AnimMap["E"] = monkey2ESprites
	monkey2AnimMap["W"] = monkey2WSprites
	monkey2AnimMap["S"] = monkey2SSprites
	monkey3AnimMap := make(map[string][]*pixel.Sprite)
	monkey3AnimMap["N"] = monkey3NSprites
	monkey3AnimMap["E"] = monkey3ESprites
	monkey3AnimMap["W"] = monkey3WSprites
	monkey3AnimMap["S"] = monkey3SSprites
	monkey4AnimMap := make(map[string][]*pixel.Sprite)
	monkey4AnimMap["N"] = monkey4NSprites
	monkey4AnimMap["E"] = monkey4ESprites
	monkey4AnimMap["W"] = monkey4WSprites
	monkey4AnimMap["S"] = monkey4SSprites
	monkey5AnimMap := make(map[string][]*pixel.Sprite)
	monkey5AnimMap["N"] = monkey5NSprites
	monkey5AnimMap["E"] = monkey5ESprites
	monkey5AnimMap["W"] = monkey5WSprites
	monkey5AnimMap["S"] = monkey5SSprites
	monkey6AnimMap := make(map[string][]*pixel.Sprite)
	monkey6AnimMap["N"] = monkey6NSprites
	monkey6AnimMap["E"] = monkey6ESprites
	monkey6AnimMap["W"] = monkey6WSprites
	monkey6AnimMap["S"] = monkey6SSprites
	monkey7AnimMap := make(map[string][]*pixel.Sprite)
	monkey7AnimMap["N"] = monkey7NSprites
	monkey7AnimMap["E"] = monkey7ESprites
	monkey7AnimMap["W"] = monkey7WSprites
	monkey7AnimMap["S"] = monkey7SSprites

	return monkey0AnimMap, monkey1AnimMap, monkey2AnimMap, monkey3AnimMap, monkey4AnimMap, monkey5AnimMap, monkey6AnimMap, monkey7AnimMap
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
