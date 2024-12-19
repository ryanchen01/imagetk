package imagetk

import (
	"encoding/binary"
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

// solveLinearSystem solves the linear system Ax = b using Gaussian elimination.
func solveLinearSystem(A [][]float64, b []float64) ([]float64, error) {
	n := len(A)

	// Augment A with b
	augmented := make([][]float64, n)
	for i := range A {
		augmented[i] = append(append([]float64{}, A[i]...), b[i])
	}

	// Gaussian elimination
	for i := 0; i < n; i++ {
		// Make the diagonal element 1
		if math.Abs(augmented[i][i]) < 1e-9 {
			return nil, fmt.Errorf("matrix is singular or nearly singular")
		}

		for k := i + 1; k < n; k++ {
			ratio := augmented[k][i] / augmented[i][i]
			for j := i; j <= n; j++ {
				augmented[k][j] -= ratio * augmented[i][j]
			}
		}
	}

	// Back substitution
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := augmented[i][n]
		for j := i + 1; j < n; j++ {
			sum -= augmented[i][j] * x[j]
		}
		x[i] = sum / augmented[i][i]
	}

	return x, nil
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
	// Step 1: Compute y = x - o
	p := make([]float64, img.dimension)
	for i := 0; i < int(img.dimension); i++ {
		p[i] = point[i] - img.origin[i]
	}

	var A [][]float64
	if img.dimension == 2 {
		A = [][]float64{
			{img.direction[0] * img.spacing[0], img.direction[3] * img.spacing[1]},
			{img.direction[1] * img.spacing[0], img.direction[4] * img.spacing[1]},
		}
	} else if img.dimension == 3 {
		A = [][]float64{
			{img.direction[0] * img.spacing[0], img.direction[3] * img.spacing[1], img.direction[6] * img.spacing[2]},
			{img.direction[1] * img.spacing[0], img.direction[4] * img.spacing[1], img.direction[7] * img.spacing[2]},
			{img.direction[2] * img.spacing[0], img.direction[5] * img.spacing[1], img.direction[8] * img.spacing[2]},
		}
	}

	i_float, err := solveLinearSystem(A, p)
	if err != nil {
		return 0.0, err
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
			value := img.pixels[linearIndex]
			pixelValue = float64(uint8(value))
		case PixelTypeInt8:
			value := img.pixels[linearIndex]
			pixelValue = float64(int8(value))
		case PixelTypeUInt16:
			value := binary.LittleEndian.Uint16(img.pixels[linearIndex*2 : linearIndex*2+2])
			pixelValue = float64(uint16(value))
		case PixelTypeInt16:
			value := binary.LittleEndian.Uint16(img.pixels[linearIndex*2 : linearIndex*2+2])
			pixelValue = float64(int16(value))
		case PixelTypeUInt32:
			value := binary.LittleEndian.Uint32(img.pixels[linearIndex*4 : linearIndex*4+4])
			pixelValue = float64(uint32(value))
		case PixelTypeInt32:
			value := binary.LittleEndian.Uint32(img.pixels[linearIndex*4 : linearIndex*4+4])
			pixelValue = float64(int32(value))
		case PixelTypeUInt64:
			value := binary.LittleEndian.Uint64(img.pixels[linearIndex*8 : linearIndex*8+8])
			pixelValue = float64(uint64(value))
		case PixelTypeInt64:
			value := binary.LittleEndian.Uint64(img.pixels[linearIndex*8 : linearIndex*8+8])
			pixelValue = float64(int64(value))
		case PixelTypeFloat32:
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[linearIndex*4 : linearIndex*4+4]))
			pixelValue = float64(value)
		case PixelTypeFloat64:
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[linearIndex*8 : linearIndex*8+8]))
			pixelValue = value
		default:
			return 0.0, fmt.Errorf("unsupported pixel type")
		}

		interpolatedValue += pixelValue * weight
	}

	return interpolatedValue, nil
}
