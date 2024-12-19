package imagetk

import (
	"fmt"
	"runtime"
	"sync"
)

type LinearInterpolator struct {
	Size      []uint32
	Spacing   []float64
	Origin    []float64
	Direction [9]float64
	FillType  int
}

type NearestInterpolator struct {
	Size      []uint32
	Spacing   []float64
	Origin    []float64
	Direction [9]float64
	FillType  int
}

const (
	FillTypeZero = iota
	FillTypeNearest
)

// Resample resamples the image using the specified interpolator.
//
// Parameters:
//   - interpolator: Interpolator
//
// Returns:
//   - *Image: The resampled image
//   - error: Error if resampling fails
func (img *Image) Resample(interpolator any) (*Image, error) {
	switch interpolator := interpolator.(type) {
	case LinearInterpolator:
		return linearResample(img, interpolator)
	case NearestInterpolator:
		return nearestResample(img, interpolator)
	default:
		return nil, fmt.Errorf("unknown interpolation type")
	}
}

func linearResample(img *Image, interpolator LinearInterpolator) (*Image, error) {
	// Create a new image
	newImg, err := NewImage(interpolator.Size, img.GetPixelType())
	if err != nil {
		return nil, err
	}

	// Copy the origin, spacing and direction with default values if not specified
	if interpolator.Origin != nil {
		newImg.SetOrigin(interpolator.Origin)
	} else {
		return nil, fmt.Errorf("origin is not specified")
	}

	if interpolator.Spacing != nil {
		newImg.SetSpacing(interpolator.Spacing)
	} else {
		return nil, fmt.Errorf("spacing is not specified")
	}

	if interpolator.Direction != [9]float64{} {
		newImg.SetDirection(interpolator.Direction)
	} else {
		return nil, fmt.Errorf("direction is not specified")
	}

	if interpolator.Size == nil {
		return nil, fmt.Errorf("size is not specified")
	}

	numPixels := 1
	for i := 0; i < len(interpolator.Size); i++ {
		numPixels *= int(interpolator.Size[i])
	}

	strides := make([]int, len(interpolator.Size))
	strides[len(interpolator.Size)-1] = 1
	for i := len(interpolator.Size) - 2; i >= 0; i-- {
		strides[i] = strides[i+1] * int(interpolator.Size[i+1])
	}

	numGoroutines := uint32(runtime.NumCPU())
	chunkSize := (uint32(numPixels) + numGoroutines - 1) / numGoroutines

	wg := sync.WaitGroup{}
	for chunk := uint32(0); chunk < numGoroutines; chunk++ {
		start := chunk * chunkSize
		end := start + chunkSize
		if end > uint32(numPixels) {
			end = uint32(numPixels)
		}
		wg.Add(1)
		go func(start, end uint32) {
			defer wg.Done()
			for i := start; i < end; i++ {
				physicalPoint := make([]float64, len(interpolator.Size))
				idx := i
				indices := make([]float64, len(interpolator.Size))
				for j := 0; j < len(interpolator.Size); j++ {
					indices[j] = float64(idx / uint32(strides[j]))
					idx %= uint32(strides[j])
				}

				// Apply direction matrix correctly
				for j := 0; j < len(interpolator.Size); j++ { // physical dimensions
					for k := 0; k < len(interpolator.Size); k++ { // image dimensions
						physicalPoint[j] += interpolator.Direction[j*3+k] * indices[k]
					}
					physicalPoint[j] = physicalPoint[j]*interpolator.Spacing[j] + interpolator.Origin[j]
				}

				value, err := img.GetPixelFromPoint(physicalPoint, interpolator.FillType)
				if err != nil {
					return
				}

				pixelValue, err := getValueAsPixelType(value, newImg.pixelType)
				if err != nil {
					return
				}
				index := make([]uint32, len(interpolator.Size))
				idx = i
				for j := 0; j < len(interpolator.Size); j++ {
					index[j] = uint32(idx / uint32(strides[j]))
					idx %= uint32(strides[j])
				}
				err = newImg.SetPixel(index, pixelValue)
				if err != nil {
					return
				}
			}
		}(start, end)
	}
	wg.Wait()

	return newImg, nil
}

func nearestResample(img *Image, interpolator NearestInterpolator) (*Image, error) {
	// Create a new image
	newImg, err := NewImage(interpolator.Size, img.GetPixelType())
	if err != nil {
		return nil, err
	}

	// Copy the origin, spacing and direction with default values if not specified
	if interpolator.Origin != nil {
		newImg.SetOrigin(interpolator.Origin)
	} else {
		return nil, fmt.Errorf("origin is not specified")
	}

	if interpolator.Spacing != nil {
		newImg.SetSpacing(interpolator.Spacing)
	} else {
		return nil, fmt.Errorf("spacing is not specified")
	}

	if interpolator.Direction != [9]float64{} {
		newImg.SetDirection(interpolator.Direction)
	} else {
		return nil, fmt.Errorf("direction is not specified")
	}

	if interpolator.Size == nil {
		return nil, fmt.Errorf("size is not specified")
	}

	numPixels := 1
	for i := 0; i < len(interpolator.Size); i++ {
		numPixels *= int(interpolator.Size[i])
	}

	strides := make([]int, len(interpolator.Size))
	strides[len(interpolator.Size)-1] = 1
	for i := len(interpolator.Size) - 2; i >= 0; i-- {
		strides[i] = strides[i+1] * int(interpolator.Size[i+1])
	}

	numGoroutines := uint32(runtime.NumCPU())
	chunkSize := (uint32(numPixels) + numGoroutines - 1) / numGoroutines

	wg := sync.WaitGroup{}
	for chunk := uint32(0); chunk < numGoroutines; chunk++ {
		start := chunk * chunkSize
		end := start + chunkSize
		if end > uint32(numPixels) {
			end = uint32(numPixels)
		}
		wg.Add(1)
		go func(start, end uint32) {
			defer wg.Done()
			for i := start; i < end; i++ {
				physicalPoint := make([]float64, len(interpolator.Size))
				idx := i
				indices := make([]float64, len(interpolator.Size))
				for j := 0; j < len(interpolator.Size); j++ {
					indices[j] = float64(idx / uint32(strides[j]))
					idx %= uint32(strides[j])
				}

				// Apply direction matrix correctly
				for j := 0; j < len(interpolator.Size); j++ { // physical dimensions
					for k := 0; k < len(interpolator.Size); k++ { // image dimensions
						physicalPoint[j] += interpolator.Direction[j*3+k] * indices[k]
					}
					physicalPoint[j] = physicalPoint[j]*interpolator.Spacing[j] + interpolator.Origin[j]
				}

				// Transform the physical point back to input image space
				inputPoint := make([]float64, len(interpolator.Size))
				for j := range physicalPoint {
					// Convert physical point to input image space using correct transformation
					// Subtract 0.5 because pixel coordinates are at center of pixel
					inputPoint[j] = ((physicalPoint[j] - img.origin[j]) / img.spacing[j]) - 0.5
					// Round to nearest integer for nearest neighbor
					inputPoint[j] = float64(int(inputPoint[j]+0.5)) + 0.5
				}

				value, err := img.GetPixelFromPoint(inputPoint, FillTypeNearest)
				if err != nil {
					return
				}

				pixelValue, err := getValueAsPixelType(value, newImg.pixelType)
				if err != nil {
					return
				}

				index := make([]uint32, len(interpolator.Size))
				idx = i
				for j := 0; j < len(interpolator.Size); j++ {
					index[j] = uint32(idx / uint32(strides[j]))
					idx %= uint32(strides[j])
				}

				err = newImg.SetPixel(index, pixelValue)
				if err != nil {
					return
				}
			}
		}(start, end)
	}
	wg.Wait()

	return newImg, nil
}
