package imagetk

import "fmt"

type LinearInterpolation struct {
	size      []uint32
	spacing   []float64
	origin    []float64
	direction [9]float64
}

func (img *Image) Resample(interpolation any) (*Image, error) {
	switch interpolation := interpolation.(type) {
	case LinearInterpolation:
		return linearResample(img, interpolation)
	default:
		return nil, fmt.Errorf("unknown interpolation type")
	}
}

func linearResample(img *Image, interpolation LinearInterpolation) (*Image, error) {
	// Create a new image
	newImg, err := NewImage(interpolation.size, img.GetPixelType())
	if err != nil {
		return nil, err
	}

	// Copy the origin, spacing and direction
	newImg.SetOrigin(interpolation.origin)
	newImg.SetSpacing(interpolation.spacing)
	newImg.SetDirection(interpolation.direction)

	numPixels := 1
	for i := 0; i < len(interpolation.size); i++ {
		numPixels *= int(interpolation.size[i])
	}

	for i := 0; i < numPixels; i++ {
		point := make([]float64, len(interpolation.size))
		for j := 0; j < len(interpolation.size); j++ {
			point[j] = interpolation.origin[j] + float64(i)*interpolation.spacing[j]
		}
		value, err := img.GetPixelFromPoint(point)
		if err != nil {
			return nil, err
		}
		fmt.Println(point, value)
		pixelValue, err := getValueAsPixelType(value, newImg.pixelType)
		if err != nil {
			return nil, err
		}
		index := make([]uint32, len(interpolation.size))
		idx := i
		for j := 0; j < len(interpolation.size); j++ {
			index[j] = uint32(idx % int(interpolation.size[j]))
			idx /= int(interpolation.size[j])
		}
		err = newImg.SetPixel(index, pixelValue)
		if err != nil {
			return nil, err
		}
	}

	return newImg, nil
}
