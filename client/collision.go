package client

import (
	"github.com/faiface/pixel"
)

var (
	collisionRs []pixel.Rect

	debugOverride bool
)

func rectCollides(r pixel.Rect) bool {
	for _, col := range collisionRs {
		if col.Intersect(r) != pixel.ZR {
			// fmt.Println("Collision:", col)
			// fmt.Println("Collision:", r)
			return true
		}
	}
	return false
}
