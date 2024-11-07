package imagetk

import (
	"fmt"
	"math"
)

func (img *Image) GetPixelFromPoint(point []float64, fillType int) (float64, error) {
	if len(point) != int(img.dimension) {
		return 0.0, fmt.Errorf("point dimension does not match image dimension")
	}

	dim := int(img.dimension)
	index := make([]int64, dim)
	t := make([]float64, dim)
	isExact := true

	for i := 0; i < dim; i++ {
		// Compute the float index
		floatIndex := (point[i] - img.origin[i]) / img.spacing[i]
		index[i] = int64(math.Floor(floatIndex))
		t[i] = floatIndex - float64(index[i])

		if t[i] != 0 {
			isExact = false
		}

		// Check index bounds
		if fillType == FillTypeZero {
			if index[i] < -1 || index[i] >= int64(img.size[i]) {
				return 0.0, nil
			}
		} else if fillType == FillTypeNearest {
			if index[i] < 0 {
				index[i] = 0
			} else if index[i] >= int64(img.size[i]) {
				index[i] = int64(img.size[i]) - 1
			}
		} else {
			return 0.0, fmt.Errorf("unknown fill type")
		}
	}
	// If the point is exactly on a pixel, return that pixel value
	if isExact {
		uintIndex := make([]uint32, dim)
		for i := 0; i < dim; i++ {
			uintIndex[i] = uint32(index[i])
		}
		pixelValue, err := img.GetPixelAsFloat64(uintIndex)
		if err != nil {
			return 0.0, err
		}
		return pixelValue, nil
	}

	// Compute the corners and weights for interpolation
	numCorners := 1 << dim // 2^dim combinations
	pixelValue := 0.0
	for i := 0; i < numCorners; i++ {
		weight := 1.0
		cornerIndex := make([]int64, dim)
		validCorner := true
		for j := 0; j < dim; j++ {
			if (i>>j)&1 == 0 {
				// Use the floor index
				cornerIndex[j] = index[j]
				weight *= (1 - t[j])
			} else {
				// Use the next index
				cornerIndex[j] = index[j] + 1
				weight *= t[j]
			}

			// Check cornerIndex bounds
			if cornerIndex[j] < 0 || cornerIndex[j] >= int64(img.size[j]) {
				validCorner = false
				break
			}
		}

		if !validCorner || weight == 0 {
			continue
		}

		// Convert cornerIndex to []uint32 for GetPixelAsFloat64
		uintCornerIndex := make([]uint32, dim)
		for j := 0; j < dim; j++ {
			uintCornerIndex[j] = uint32(cornerIndex[j])
		}

		cornerValue, err := img.GetPixelAsFloat64(uintCornerIndex)
		if err != nil {
			return 0.0, err
		}
		pixelValue += weight * cornerValue
	}

	return pixelValue, nil
}
