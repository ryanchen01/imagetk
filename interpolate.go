package imagetk

import "fmt"

type LinearInterpolation struct {
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
	newImg, err := NewImage(interpolation.Size, img.GetPixelType())
	if err != nil {
		return nil, err
	}

	// Copy the origin, spacing and direction
	newImg.SetOrigin(interpolation.Origin)
	newImg.SetSpacing(interpolation.Spacing)
	newImg.SetDirection(interpolation.Direction)

	numPixels := 1
	for i := 0; i < len(interpolation.Size); i++ {
		numPixels *= int(interpolation.Size[i])
	}

	strides := make([]int, len(interpolation.Size))
	strides[len(interpolation.Size)-1] = 1
	for i := len(interpolation.Size) - 2; i >= 0; i-- {
		strides[i] = strides[i+1] * int(interpolation.Size[i+1])
	}

	for i := 0; i < numPixels; i++ {
		point := make([]float64, len(interpolation.Size))
		idx := i
		for j := 0; j < len(interpolation.Size); j++ {
			point[j] = float64(idx/strides[j])*interpolation.Spacing[j] + interpolation.Origin[j]
			idx %= strides[j]
		}
		value, err := img.GetPixelFromPoint(point, interpolation.FillType)
		if err != nil {
			return nil, err
		}

		pixelValue, err := getValueAsPixelType(value, newImg.pixelType)
		if err != nil {
			return nil, err
		}
		index := make([]uint32, len(interpolation.Size))
		idx = i
		for j := 0; j < len(interpolation.Size); j++ {
			index[j] = uint32(idx / strides[j])
			idx %= strides[j]
		}
		err = newImg.SetPixel(index, pixelValue)
		if err != nil {
			return nil, err
		}
	}

	return newImg, nil
}
