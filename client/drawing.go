package client

import (
	"image"
	"os"

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
