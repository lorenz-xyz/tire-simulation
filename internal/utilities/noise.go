package utilities

import (
	"fmt"
	"math"

	noise "github.com/cjslep/noise"
)

func GenerateNoise() {
	perlinGenerator := noise.NewPerlin(1)
	var x float64 = 0.5
	for y := 0.0; y < 1; y += 0.1 {
		//noiseVal := perlinGenerator.Noise(x, y)
		noiseVal := perlinGenerator.Noise(x+perlinGenerator.Noise(x, y), y+perlinGenerator.Noise(x, y))
		//fmt.Println("x:", x, "y:", y, noiseVal)

		modVal := 1 - math.Abs(noiseVal)
		cutVal := int(modVal * 50)

		for i := 0; i < cutVal; i++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
