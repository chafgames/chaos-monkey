package client

import (
	"fmt"

	"github.com/faiface/pixel"
)

var (
	collisionRs []pixel.Rect

	debugOverride bool
)

func rectCollides(r pixel.Rect) bool {
	for _, col := range collisionRs {
		if col.Intersect(r) != pixel.ZR {
			fmt.Println(col, r)
			return true
		}
	}
	return false
}
