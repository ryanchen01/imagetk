package imagetk

import (
	"runtime"
	"sync"
)

const (
	MORPH_OPEN = iota
	MORPH_CLOSE
)

// BinaryDilate dilates the binary image.
// Parameters:
//   - image: The image to dilate.
//   - kernelSize: The size of the kernel to use for the dilation.
//
// Returns:
//   - *Image: The resulting image after dilation.
//   - error: An error if the operation fails.
func BinaryDilate(image *Image, kernelSize int) (*Image, error) {
	switch image.GetDimension() {
	case 2:
		return binaryDilate2D(image, kernelSize)
	case 3:
		return binaryDilate3D(image, kernelSize)
	default:
		return nil, nil
	}
}

// BinaryErode erodes the binary image.
// Parameters:
//   - image: The image to erode.
//   - kernelSize: The size of the kernel to use for the erosion.
//
// Returns:
//   - *Image: The resulting image after erosion.
//   - error: An error if the operation fails.
func BinaryErode(image *Image, kernelSize int) (*Image, error) {
	switch image.GetDimension() {
	case 2:
		return binaryErode2D(image, kernelSize)
	case 3:
		return binaryErode3D(image, kernelSize)
	default:
		return nil, nil
	}
}

// Morphology performs morphological operations on the image.
// Parameters:
//   - image: The image to perform the morphological operation on.
//   - operation: The morphological operation to perform.
//   - kernelSize: The size of the kernel to use for the morphological operation.
//   - iterations: The number of iterations to perform.
//
// Returns:
//   - *Image: The resulting image after the morphological operation.
//   - error: An error if the operation fails.
func Morphology(image *Image, operation, kernelSize, iterations int) (*Image, error) {
	var output *Image
	var err error
	switch operation {
	case MORPH_OPEN:
		for i := 0; i < iterations; i++ {
			output, err = BinaryErode(image, kernelSize)
			if err != nil {
				return nil, err
			}
		}
		for i := 0; i < iterations; i++ {
			output, err = BinaryDilate(output, kernelSize)
			if err != nil {
				return nil, err
			}
		}
	case MORPH_CLOSE:
		for i := 0; i < iterations; i++ {
			output, err = BinaryDilate(image, kernelSize)
			if err != nil {
				return nil, err
			}
		}
		for i := 0; i < iterations; i++ {
			output, err = BinaryErode(output, kernelSize)
			if err != nil {
				return nil, err
			}
		}
	}
	return output, nil
}

func binaryDilate3D(image *Image, kernelSize int) (*Image, error) {
	size := image.GetSize()
	int8Array := make([][][]int8, size[2])
	expandedArray := make([][][]int8, size[2])
	for z := uint32(0); z < size[2]; z++ {
		int8Array[z] = make([][]int8, size[1])
		expandedArray[z] = make([][]int8, size[1])
		for y := uint32(0); y < size[1]; y++ {
			int8Array[z][y] = make([]int8, size[0])
			expandedArray[z][y] = make([]int8, size[0])
			for x := uint32(0); x < size[0]; x++ {
				val, err := image.GetPixelAsInt8([]uint32{x, y, z})
				if err != nil {
					return nil, err
				}
				if val > 0 {
					int8Array[z][y][x] = 1
				} else {
					int8Array[z][y][x] = 0
				}
			}
		}
	}

	numGoroutines := uint32(runtime.NumCPU())
	chunkSize := size[2] / numGoroutines
	if chunkSize*numGoroutines < size[2] {
		chunkSize += 1
	}
	wg := sync.WaitGroup{}
	for chunk := uint32(0); chunk < numGoroutines; chunk++ {
		start := chunk * chunkSize
		end := start + chunkSize
		if end > size[2] {
			end = size[2]
		}
		wg.Add(1)
		go func(start, end uint32) {
			defer wg.Done()
			for z := start; z < end; z++ {
				for y := uint32(0); y < size[1]; y++ {
					for x := uint32(0); x < size[0]; x++ {
						// Check if the current pixel is set to 1
						if int8Array[z][y][x] == 1 {
							// Expand around the pixel, respecting boundaries
							for i := -kernelSize / 2; i <= kernelSize/2; i++ {
								for j := -kernelSize / 2; j <= kernelSize/2; j++ {
									for k := -kernelSize / 2; k <= kernelSize/2; k++ {
										newZ := int(z) + i
										newY := int(y) + j
										newX := int(x) + k

										// Ensure the new coordinates are within bounds
										if newZ >= 0 && newZ < int(size[2]) && newY >= 0 && newY < int(size[1]) && newX >= 0 && newX < int(size[0]) {
											expandedArray[newZ][newY][newX] = 1
										}
									}
								}
							}
						}
					}
				}
			}
		}(start, end)
	}
	wg.Wait()

	newImage, err := GetImageFromArray(expandedArray)
	if err != nil {
		return nil, err
	}

	newImage.SetOrigin(image.GetOrigin())
	newImage.SetSpacing(image.GetSpacing())
	newImage.SetDirection(image.GetDirection())
	return newImage, nil
}

func binaryDilate2D(image *Image, kernelSize int) (*Image, error) {
	size := image.GetSize()
	int8Array := make([][]int8, size[1])
	expandedArray := make([][]int8, size[1])
	for y := uint32(0); y < size[1]; y++ {
		int8Array[y] = make([]int8, size[0])
		expandedArray[y] = make([]int8, size[0])
		for x := uint32(0); x < size[0]; x++ {
			val, err := image.GetPixelAsInt8([]uint32{x, y})
			if err != nil {
				return nil, err
			}
			if val > 0 {
				int8Array[y][x] = 1
			} else {
				int8Array[y][x] = 0
			}
		}
	}
	for y := uint32(0); y < size[1]; y++ {
		for x := uint32(0); x < size[0]; x++ {
			// Check if the current pixel is set to 1
			if int8Array[y][x] == 1 {
				// Expand around the pixel, respecting boundaries
				for j := -kernelSize / 2; j <= kernelSize/2; j++ {
					for k := -kernelSize / 2; k <= kernelSize/2; k++ {
						newY := int(y) + j
						newX := int(x) + k

						// Ensure the new coordinates are within bounds
						if newY >= 0 && newY < int(size[1]) && newX >= 0 && newX < int(size[0]) {
							expandedArray[newY][newX] = 1
						}
					}
				}
			}
		}
	}

	newImage, err := GetImageFromArray(expandedArray)
	if err != nil {
		return nil, err
	}

	newImage.SetOrigin(image.GetOrigin())
	newImage.SetSpacing(image.GetSpacing())
	newImage.SetDirection(image.GetDirection())
	return newImage, nil
}

func binaryErode3D(image *Image, kernelSize int) (*Image, error) {
	size := image.GetSize()
	int8Array := make([][][]int8, size[2])
	erodedArray := make([][][]int8, size[2])

	// Initialize arrays and copy input
	for z := uint32(0); z < size[2]; z++ {
		int8Array[z] = make([][]int8, size[1])
		erodedArray[z] = make([][]int8, size[1])
		for y := uint32(0); y < size[1]; y++ {
			int8Array[z][y] = make([]int8, size[0])
			erodedArray[z][y] = make([]int8, size[0])
			for x := uint32(0); x < size[0]; x++ {
				val, err := image.GetPixelAsInt8([]uint32{x, y, z})
				if err != nil {
					return nil, err
				}
				if val > 0 {
					int8Array[z][y][x] = 1
					erodedArray[z][y][x] = 1
				}
			}
		}
	}

	halfKernel := kernelSize / 2

	numGoroutines := uint32(runtime.NumCPU())
	chunkSize := size[2] / numGoroutines
	if chunkSize*numGoroutines < size[2] {
		chunkSize += 1
	}
	wg := sync.WaitGroup{}
	for chunk := uint32(0); chunk < numGoroutines; chunk++ {
		start := chunk * chunkSize
		end := start + chunkSize
		if end > size[2] {
			end = size[2]
		}
		wg.Add(1)
		go func(start, end uint32) {
			defer wg.Done()
			for z := start; z < end; z++ {
				for y := uint32(0); y < size[1]; y++ {
					for x := uint32(0); x < size[0]; x++ {
						// Only process pixels that are 1 in the input
						if int8Array[z][y][x] == 1 {
							// Check all pixels in the kernel neighborhood
							for dz := -halfKernel; dz <= halfKernel; dz++ {
								for dy := -halfKernel; dy <= halfKernel; dy++ {
									for dx := -halfKernel; dx <= halfKernel; dx++ {
										newZ := int(z) + dz
										newY := int(y) + dy
										newX := int(x) + dx

										// If any neighbor is outside bounds or 0, erode the current pixel
										if newZ < 0 || newZ >= int(size[2]) ||
											newY < 0 || newY >= int(size[1]) ||
											newX < 0 || newX >= int(size[0]) ||
											int8Array[newZ][newY][newX] == 0 {
											erodedArray[z][y][x] = 0
											goto nextPixel // Break out of all loops
										}
									}
								}
							}
						}
					nextPixel:
					}
				}
			}
		}(start, end)
	}
	wg.Wait()

	newImage, err := GetImageFromArray(erodedArray)
	if err != nil {
		return nil, err
	}

	newImage.SetOrigin(image.GetOrigin())
	newImage.SetSpacing(image.GetSpacing())
	newImage.SetDirection(image.GetDirection())
	return newImage, nil
}

func binaryErode2D(image *Image, kernelSize int) (*Image, error) {
	size := image.GetSize()
	int8Array := make([][]int8, size[1])
	erodedArray := make([][]int8, size[1])

	// Initialize arrays and copy input
	for y := uint32(0); y < size[1]; y++ {
		int8Array[y] = make([]int8, size[0])
		erodedArray[y] = make([]int8, size[0])
		for x := uint32(0); x < size[0]; x++ {
			val, err := image.GetPixelAsInt8([]uint32{x, y})
			if err != nil {
				return nil, err
			}
			if val > 0 {
				int8Array[y][x] = 1
				erodedArray[y][x] = 1
			}
		}
	}

	halfKernel := kernelSize / 2

	for y := uint32(0); y < size[1]; y++ {
		for x := uint32(0); x < size[0]; x++ {
			// Only process pixels that are 1 in the input
			if int8Array[y][x] == 1 {
				// Check all pixels in the kernel neighborhood
				for dy := -halfKernel; dy <= halfKernel; dy++ {
					for dx := -halfKernel; dx <= halfKernel; dx++ {
						newY := int(y) + dy
						newX := int(x) + dx

						// If any neighbor is outside bounds or 0, erode the current pixel
						if newY < 0 || newY >= int(size[1]) ||
							newX < 0 || newX >= int(size[0]) ||
							int8Array[newY][newX] == 0 {
							erodedArray[y][x] = 0
							goto nextPixel
						}
					}
				}
			}
		nextPixel:
		}
	}

	newImage, err := GetImageFromArray(erodedArray)
	if err != nil {
		return nil, err
	}

	newImage.SetOrigin(image.GetOrigin())
	newImage.SetSpacing(image.GetSpacing())
	newImage.SetDirection(image.GetDirection())
	return newImage, nil
}
