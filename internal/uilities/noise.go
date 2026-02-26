package utilities

import (
	noise "github.com/cjslep/noise"
	"fmt"
)

func GenerateNoise {
	perlinGenerator := noise.NewPerlin(1)
	val := perlinGenerator.Noise(0.5, 0.5)
	fmt.Println(val)
}
