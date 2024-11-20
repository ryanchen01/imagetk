# ImageTK
[![tests](https://github.com/ryanchen01/imagetk/actions/workflows/tests.yml/badge.svg)](https://github.com/ryanchen01/imagetk/actions/workflows/tests.yml)

ImageTK is a Go library for handling and manipulating multi-dimensional images with support for various pixel types and image formats.

## Features

- Support for multiple pixel types (uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64)
- 2D and 3D image handling
- Image resampling
- Raw and MetaImage (MHD) file format support
- Basic image statistics (min, max, mean, median, std)
- Physical space transformations

## Installation

```bash
go get github.com/ryanchen01/imagetk
```

## Usage

### Creating a New Image

```go
// Create a 2D image with float32 pixels
size := []uint32{10, 10}
img, err := NewImage(size, PixelTypeFloat32)
if err != nil {
    log.Fatal(err)
}
```

### Setting and Getting Pixel Values

```go
// Set a pixel value
err = img.SetPixel([]uint32{0, 0}, float32(1.0))
if err != nil {
    log.Fatal(err)
}

// Get a pixel value
value, err := img.GetPixel([]uint32{0, 0})
if err != nil {
    log.Fatal(err)
}
```

### Image Resampling

```go
interpolator := LinearInterpolator{
    Size:      []uint32{20, 20},
    Spacing:   []float64{0.5, 0.5},
    Origin:    []float64{0.0, 0.0},
    Direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
}

resampledImg, err := img.Resample(interpolator)
if err != nil {
    log.Fatal(err)
}
```

### Reading and Writing Images

```go
// Read an MHD image
img, err := ReadImage("image.mhd", ImageTypeMHD, nil)
if err != nil {
    log.Fatal(err)
}

// Write an MHD image
err = WriteImage(img, "output.mhd", ImageTypeMHD)
if err != nil {
    log.Fatal(err)
}
```

## Supported File Formats

- Raw binary files
- MetaImage format (MHD/RAW pairs)

## Image Properties

- Dimension (2D or 3D)
- Size (number of pixels in each dimension)
- Pixel spacing (voxel size in each dimension)
- Origin (physical coordinate of the first pixel)
- Direction cosines (orientation of the image)
- Pixel type (uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## References

The implementation includes code for:
- Image creation and manipulation (see `image.go`)
- File I/O operations (see `io.go`)
- Interpolation (see `interpolate.go`)
- Statistical operations (see `stats.go`)
- Physical point transformations (see `point.go`)

For detailed implementation examples, please refer to the test files in the repository.

## TODO
- [ ] Cubic interpolation
- [ ] DICOM file format support
- [ ] NRRD file format support
- [x] Dilate/Erode/Open/Close morphological operations
