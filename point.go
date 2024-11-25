package imagetk

import (
	"fmt"
	"math"
)

// invert2x2 computes the inverse of a 2x2 matrix.
func invert2x2(m [4]float64) ([4]float64, error) {
	a, b := m[0], m[1]
	c, d := m[2], m[3]
	det := a*d - b*c
	if det == 0 {
		return [4]float64{}, fmt.Errorf("direction matrix is singular")
	}
	invDet := 1.0 / det
	return [4]float64{
		d * invDet,
		-b * invDet,
		-c * invDet,
		a * invDet,
	}, nil
}

// invert3x3 computes the inverse of a 3x3 matrix.
func invert3x3(m [9]float64) ([9]float64, error) {
	a, b, c := m[0], m[1], m[2]
	d, e, f := m[3], m[4], m[5]
	g, h, i := m[6], m[7], m[8]

	det := a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
	if det == 0 {
		return [9]float64{}, fmt.Errorf("direction matrix is singular")
	}
	invDet := 1.0 / det

	return [9]float64{
		(e*i - f*h) * invDet,
		-(b*i - c*h) * invDet,
		(b*f - c*e) * invDet,
		-(d*i - f*g) * invDet,
		(a*i - c*g) * invDet,
		-(a*f - c*d) * invDet,
		(d*h - e*g) * invDet,
		-(a*h - b*g) * invDet,
		(a*e - b*d) * invDet,
	}, nil
}

// GetPixelFromPoint returns the interpolated pixel value at a given physical point in the image.
// Parameters:
//   - point: A slice of float64 representing the physical point.
//   - fillType: The type of fill to use if the point is outside the image bounds.
//
// Returns:
//   - float64: The interpolated pixel value.
//   - error: An error if the operation fails.
func (img *Image) GetPixelFromPoint(point []float64, fillType int) (float64, error) {
	if len(point) != int(img.dimension) {
		return 0.0, fmt.Errorf("point dimension does not match image dimension")
	}

	// Step 1: Compute y = x - o
	y := make([]float64, img.dimension)
	for i := 0; i < int(img.dimension); i++ {
		y[i] = point[i] - img.origin[i]
	}

	// Step 2: Compute D_inv
	var D_inv []float64
	if img.dimension == 2 {
		D := [4]float64{
			img.direction[0], img.direction[1],
			img.direction[3], img.direction[4],
		}
		inv, err := invert2x2(D)
		if err != nil {
			return 0.0, err
		}
		D_inv = inv[:]
	} else if img.dimension == 3 {
		D := [9]float64{
			img.direction[0], img.direction[1], img.direction[2],
			img.direction[3], img.direction[4], img.direction[5],
			img.direction[6], img.direction[7], img.direction[8],
		}
		inv, err := invert3x3(D)
		if err != nil {
			return 0.0, err
		}
		D_inv = inv[:]
	} else {
		return 0.0, fmt.Errorf("unsupported dimension: %d", img.dimension)
	}

	// Step 3: Compute p = D_inv * y
	p := make([]float64, img.dimension)
	if img.dimension == 2 {
		p[0] = D_inv[0]*y[0] + D_inv[1]*y[1]
		p[1] = D_inv[2]*y[0] + D_inv[3]*y[1]
	} else {
		p[0] = D_inv[0]*y[0] + D_inv[1]*y[1] + D_inv[2]*y[2]
		p[1] = D_inv[3]*y[0] + D_inv[4]*y[1] + D_inv[5]*y[2]
		p[2] = D_inv[6]*y[0] + D_inv[7]*y[1] + D_inv[8]*y[2]
	}

	// Step 4: Compute i_float = p / s
	i_float := make([]float64, img.dimension)
	for i := 0; i < int(img.dimension); i++ {
		if img.spacing[i] == 0 {
			return 0.0, fmt.Errorf("spacing[%d] is zero", i)
		}
		i_float[i] = p[i] / img.spacing[i]
	}

	// Step 5: Compute floor and ceil indices for interpolation
	i0 := make([]int, img.dimension) // Floor indices
	i1 := make([]int, img.dimension) // Ceil indices
	weights := make([]float64, img.dimension)

	for i := 0; i < int(img.dimension); i++ {
		i0[i] = int(math.Floor(i_float[i]))
		i1[i] = i0[i] + 1
		weights[i] = i_float[i] - float64(i0[i])
	}

	// Step 6: Collect corner indices and weights
	type indexWeight struct {
		indices []int
		weight  float64
	}

	var combinations []indexWeight

	if img.dimension == 2 {
		combinations = []indexWeight{
			{indices: []int{i0[0], i0[1]}, weight: (1 - weights[0]) * (1 - weights[1])},
			{indices: []int{i1[0], i0[1]}, weight: weights[0] * (1 - weights[1])},
			{indices: []int{i0[0], i1[1]}, weight: (1 - weights[0]) * weights[1]},
			{indices: []int{i1[0], i1[1]}, weight: weights[0] * weights[1]},
		}
	} else if img.dimension == 3 {
		combinations = []indexWeight{
			{indices: []int{i0[0], i0[1], i0[2]}, weight: (1 - weights[0]) * (1 - weights[1]) * (1 - weights[2])},
			{indices: []int{i1[0], i0[1], i0[2]}, weight: weights[0] * (1 - weights[1]) * (1 - weights[2])},
			{indices: []int{i0[0], i1[1], i0[2]}, weight: (1 - weights[0]) * weights[1] * (1 - weights[2])},
			{indices: []int{i1[0], i1[1], i0[2]}, weight: weights[0] * weights[1] * (1 - weights[2])},
			{indices: []int{i0[0], i0[1], i1[2]}, weight: (1 - weights[0]) * (1 - weights[1]) * weights[2]},
			{indices: []int{i1[0], i0[1], i1[2]}, weight: weights[0] * (1 - weights[1]) * weights[2]},
			{indices: []int{i0[0], i1[1], i1[2]}, weight: (1 - weights[0]) * weights[1] * weights[2]},
			{indices: []int{i1[0], i1[1], i1[2]}, weight: weights[0] * weights[1] * weights[2]},
		}
	} else {
		return 0.0, fmt.Errorf("unsupported dimension: %d", img.dimension)
	}

	// Step 7: Interpolate
	var interpolatedValue float64
	for _, combo := range combinations {
		indices := combo.indices
		weight := combo.weight

		// Check bounds and adjust indices based on fillType
		outOfBounds := false
		for i := 0; i < int(img.dimension); i++ {
			if indices[i] < 0 || indices[i] >= int(img.size[i]) {
				if fillType == FillTypeZero {
					outOfBounds = true
					break
				} else if fillType == FillTypeNearest {
					if indices[i] < 0 {
						indices[i] = 0
					} else if indices[i] >= int(img.size[i]) {
						indices[i] = int(img.size[i]) - 1
					}
				} else {
					return 0.0, fmt.Errorf("unsupported fillType: %d", fillType)
				}
			}
		}

		if outOfBounds {
			continue // Skip this corner as it contributes zero
		}

		// Compute linear index
		linearIndex := indices[0]
		if img.dimension >= 2 {
			linearIndex += indices[1] * int(img.size[0])
		}
		if img.dimension == 3 {
			linearIndex += indices[2] * int(img.size[0]*img.size[1])
		}

		// Retrieve the pixel value
		var pixelValue float64
		switch img.pixelType {
		case PixelTypeUInt8:
			data, ok := img.pixels.([]uint8)
			if !ok || linearIndex >= len(data) {
				continue // Skip invalid data
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeInt8:
			data, ok := img.pixels.([]int8)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeUInt16:
			data, ok := img.pixels.([]uint16)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeInt16:
			data, ok := img.pixels.([]int16)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeUInt32:
			data, ok := img.pixels.([]uint32)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeInt32:
			data, ok := img.pixels.([]int32)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeUInt64:
			data, ok := img.pixels.([]uint64)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeInt64:
			data, ok := img.pixels.([]int64)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeFloat32:
			data, ok := img.pixels.([]float32)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = float64(data[linearIndex])
		case PixelTypeFloat64:
			data, ok := img.pixels.([]float64)
			if !ok || linearIndex >= len(data) {
				continue
			}
			pixelValue = data[linearIndex]
		default:
			return 0.0, fmt.Errorf("unsupported pixel type")
		}

		interpolatedValue += pixelValue * weight
	}

	return interpolatedValue, nil
}
