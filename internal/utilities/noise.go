package utilities

import (
	"fmt"

	noise "github.com/cjslep/noise"
)

func GenerateNoise() {
	perlinGenerator := noise.NewPerlin(1)
	var x float64 = 0.5
	for y := 0.0; y < 1; y += 0.02 {
		//noiseVal := perlinGenerator.Noise(x, y)
		noiseVal := generatePerlinNoise(x, y, perlinGenerator, 7)
		//fmt.Println("x:", x, "y:", y, noiseVal)

		//modVal := 1 - math.Abs(noiseVal)
		modVal := noiseVal + 1
		cutVal := int(modVal * 20)

		for i := 0; i < cutVal; i++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}

func generatePerlinNoise(x float64, y float64, generator *noise.Perlin, interator int) float64 {
	if interator > 0 {
		return generator.Noise(x+generatePerlinNoise(x, y, generator, interator-1), y+generatePerlinNoise(x, y, generator, interator-1))
	} else {
		return generator.Noise(x, y)
	}
}
