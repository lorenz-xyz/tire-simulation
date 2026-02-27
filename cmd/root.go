package cmd

import (
	"github.com/lorenz-xyz/tire-simuation/internal/render"
	"github.com/lorenz-xyz/tire-simuation/internal/utilities"
)

func Execute() {

	utilities.GenerateNoise()
	render.Test2D()
	//render.Render()

}
