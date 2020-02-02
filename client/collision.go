package client

import (
	"log"

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

func playerCollides(r pixel.Rect) bool {
	if oncallCollides(r) == true {
		return true
	} else if monkeyCollides(r) == true {
		return true
	}
	return false
}

func oncallCollides(r pixel.Rect) bool {
	if myPlayer.ID == "onhands" {
		return false
	}
	minx := myOnHands.State.IdentityMatrix[4]
	miny := myOnHands.State.IdentityMatrix[5]
	oncallRect := pixel.R(minx, miny, minx+48, miny+48)
	if myPlayer.collisionBox().Intersect(oncallRect) != pixel.ZR {
		log.Printf("%s collides with %s", myPlayer.ID, myOnHands.ID)
		return true
	}
	return false
}
func monkeyCollides(r pixel.Rect) bool {
	for _, monkey := range myMonkeys {
		if monkey.State.Active == true {
			if myPlayer.ID == monkey.ID {
				return false
			}
			minx := monkey.State.IdentityMatrix[4]
			miny := monkey.State.IdentityMatrix[5]

			monkeyRect := pixel.R(minx, miny, minx+48, miny+48)
			// _ = monkeyRect
			// monkeyVec := pixel.V(minx, miny)
			// _ = monkeyVec
			// if myPlayer.collisionBox().Contains(monkeyRect.Center()) {
			if myPlayer.collisionBox().Intersect(monkeyRect) != pixel.ZR {
				// log.Printf("COLLIDE! %+v  in vec:%+v", monkeyVec, myPlayer.collisionBox())

				log.Printf("COLLIDE! %+v  intersects with $%v", myPlayer.collisionBox(), monkeyRect)
				return true
			}
			log.Printf("NOCOLLIDE! %+v  DOESNT intersect with $%v", myPlayer.collisionBox(), monkeyRect)
			// log.Printf("%+v NOT in vec:%+v", monkeyVec, myPlayer.collisionBox())

		}
	}
	return false
}
