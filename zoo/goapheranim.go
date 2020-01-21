package zoo

import (
	"math"

	pixel "github.com/faiface/pixel"
)

type animState int

const (
	idle animState = iota
	running
	jumping
)

type gopherAnim struct {
	sheet pixel.Picture
	anims map[string][]pixel.Rect
	rate  float64

	state   animState
	counter float64
	dir     float64

	frame pixel.Rect

	sprite *pixel.Sprite
}

func (ga *gopherAnim) update(dt float64, phys *gopherPhys) {
	ga.counter += dt

	// determine the new animation state
	var newState animState
	switch {
	case !phys.ground:
		newState = jumping
	case phys.vel.Len() == 0:
		newState = idle
	case phys.vel.Len() > 0:
		newState = running
	}

	// reset the time counter if the state changed
	if ga.state != newState {
		ga.state = newState
		ga.counter = 0
	}

	// determine the correct animation frame
	switch ga.state {
	case idle:
		ga.frame = ga.anims["Front"][0]
	case running:
		i := int(math.Floor(ga.counter / ga.rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
	case jumping:
		speed := phys.vel.Y
		i := int((-speed/phys.jumpSpeed + 1) / 2 * float64(len(ga.anims["Jump"])))
		if i < 0 {
			i = 0
		}
		if i >= len(ga.anims["Jump"]) {
			i = len(ga.anims["Jump"]) - 1
		}
		ga.frame = ga.anims["Jump"][i]
	}

	// set the facing direction of the gopher
	if phys.vel.X != 0 {
		if phys.vel.X > 0 {
			ga.dir = +1
		} else {
			ga.dir = -1
		}
	}
}

func (ga *gopherAnim) draw(t pixel.Target, phys *gopherPhys) {
	if ga.sprite == nil {
		ga.sprite = pixel.NewSprite(nil, pixel.Rect{})
	}
	// draw the correct frame with the correct position and direction
	ga.sprite.Set(ga.sheet, ga.frame)
	// socketioClient.Notice(fmt.Sprintf("gopher at %d / %d", int(phys.rect.Max.X), int(phys.rect.Max.Y)))
	ga.sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			phys.rect.W()/ga.sprite.Frame().W(),
			phys.rect.H()/ga.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(-ga.dir, 1)).
		Moved(phys.rect.Center()),
	)
}
