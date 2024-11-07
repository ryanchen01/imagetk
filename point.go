package imagetk

import (
	"fmt"
	"math"
)

func (img *Image) GetPixelFromPoint(point []float64) (float64, error) {
	if len(point) != int(img.dimension) {
		return 0.0, fmt.Errorf("point dimension does not match image dimension")
	}

	index := make([]uint32, img.dimension)
	for i := 0; i < int(img.dimension); i++ {
		index[i] = uint32((point[i] - img.origin[i]) / img.spacing[i])
		if index[i] >= img.size[i] {
			return 0.0, nil
		}
	}

	corners := make([][]uint32, 1<<len(index))
	for i := 0; i < len(corners); i++ {
		corners[i] = make([]uint32, len(index))
		for j := 0; j < len(index); j++ {
			if i&(1<<j) != 0 {
				corners[i][j] = index[j] + 1
			}
		}
	}
	fmt.Println("Corners:")
	fmt.Println(corners)

	sumWeights := 0.0
	weights := make([]float64, len(corners))
	for i := 0; i < len(corners); i++ {
		distToPoint := 0.0
		for j := 0; j < len(index); j++ {
			cornerPoint := img.origin[j] + float64(corners[i][j])*img.spacing[j]
			distToPoint += math.Pow(point[j]-cornerPoint, 2)
		}
		distToPoint = math.Sqrt(distToPoint)
		weights[i] = 1 / distToPoint
		sumWeights += weights[i]
	}

	for i := 0; i < len(weights); i++ {
		weights[i] /= sumWeights
	}

	pixelValue := 0.0
	for i := 0; i < len(corners); i++ {
		cornerValue, err := img.GetPixelAsFloat64(corners[i])
		if err != nil && cornerValue != 0.0 {
			return 0.0, err
		}
		pixelValue += weights[i] * float64(cornerValue)
	}

	return pixelValue, nil
}
