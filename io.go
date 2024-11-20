package imagetk

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

const (
	ImageTypeRaw = iota
	ImageTypeMHD
)

// ReadImage reads an image from a file and returns an Image object.
// The function supports reading both raw binary files and MetaImage (MHD) files.
//
// Parameters:
//   - filename: Path to the image file to read
//   - imageType: Type of image file format
//   - pixelType: Pointer to pixel data type for raw files (can be nil for MHD files)
//
// Returns:
//   - *Image: The loaded image object
//   - error: Error if reading fails
//
// For raw files, the pixelType parameter must be specified to correctly interpret the binary data.
// For MHD files, the pixel type is determined from the header information.
func ReadImage(filename string, imageType int, pixelType *int) (*Image, error) {
	switch imageType {
	case ImageTypeRaw:
		return readImageTypeRaw(filename, *pixelType)
	case ImageTypeMHD:
		return readImageTypeMHD(filename)
	default:
		return nil, fmt.Errorf("unknown image type")
	}
}

func readImageTypeRaw(filename string, pixelType int) (*Image, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	// Determine pixel type based on file size and dimensions
	var bytesPerPixel int
	switch pixelType {
	case PixelTypeUInt8:
		bytesPerPixel = 1
	case PixelTypeInt8:
		bytesPerPixel = 1
	case PixelTypeUInt16:
		bytesPerPixel = 2
	case PixelTypeInt16:
		bytesPerPixel = 2
	case PixelTypeUInt32:
		bytesPerPixel = 4
	case PixelTypeInt32:
		bytesPerPixel = 4
	case PixelTypeUInt64:
		bytesPerPixel = 8
	case PixelTypeInt64:
		bytesPerPixel = 8
	case PixelTypeFloat32:
		bytesPerPixel = 4
	case PixelTypeFloat64:
		bytesPerPixel = 8
	}
	totalPixels := fileSize / int64(bytesPerPixel)

	// Create a new image
	img := &Image{}

	// This is a placeholder - you should modify these values based on your needs
	img.dimension = 3
	img.size = make([]uint32, img.dimension)
	xy := uint32(math.Pow(float64(totalPixels), 1.0/float64(img.dimension)))
	z := uint32(totalPixels / int64(xy*xy))
	img.size = []uint32{xy, xy, z}

	img.spacing = []float64{1.0, 1.0, 1.0}
	img.origin = []float64{0.0, 0.0, 0.0}
	img.direction = [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}

	// Set pixel type based on bytes per pixel
	switch bytesPerPixel {
	case 1:
		img.pixelType = PixelTypeUInt8
		pixels := make([]uint8, totalPixels)
		if err := binary.Read(file, binary.LittleEndian, pixels); err != nil {
			return nil, err
		}
		img.pixels = pixels
	case 2:
		img.pixelType = PixelTypeUInt16
		pixels := make([]uint16, totalPixels)
		if err := binary.Read(file, binary.LittleEndian, pixels); err != nil {
			return nil, err
		}
		img.pixels = pixels
	case 4:
		img.pixelType = PixelTypeFloat32
		pixels := make([]float32, totalPixels)
		if err := binary.Read(file, binary.LittleEndian, pixels); err != nil {
			return nil, err
		}
		img.pixels = pixels
	case 8:
		img.pixelType = PixelTypeFloat64
		pixels := make([]float64, totalPixels)
		if err := binary.Read(file, binary.LittleEndian, pixels); err != nil {
			return nil, err
		}
		img.pixels = pixels
	default:
		return nil, fmt.Errorf("unsupported bytes per pixel: %d", bytesPerPixel)
	}

	return img, nil
}

func readImageTypeMHD(filename string) (*Image, error) {
	headerFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open MHD file: %v", err)
	}
	defer headerFile.Close()

	img := &Image{}
	var rawFilename string
	var elementType string

	scanner := bufio.NewScanner(headerFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "NDims":
			fmt.Sscanf(value, "%d", &img.dimension)

		case "DimSize":
			img.size = make([]uint32, img.dimension)
			numbers := strings.Fields(value)
			for i := 0; i < len(numbers) && i < int(img.dimension); i++ {
				fmt.Sscanf(numbers[i], "%d", &img.size[i])
			}

		case "TransformMatrix":
			img.direction = [9]float64{}
			numbers := strings.Fields(value)
			for i := 0; i < len(numbers) && i < 9; i++ {
				fmt.Sscanf(numbers[i], "%f", &img.direction[i])
			}

		case "Offset":
			img.origin = make([]float64, img.dimension)
			numbers := strings.Fields(value)
			for i := 0; i < len(numbers) && i < int(img.dimension); i++ {
				fmt.Sscanf(numbers[i], "%f", &img.origin[i])
			}

		case "ElementSpacing":
			img.spacing = make([]float64, img.dimension)
			numbers := strings.Fields(value)
			for i := 0; i < len(numbers) && i < int(img.dimension); i++ {
				fmt.Sscanf(numbers[i], "%f", &img.spacing[i])
			}

		case "ElementType":
			elementType = value

		case "ElementDataFile":
			rawFilename = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading MHD file: %v", err)
	}

	// Set pixel type based on ElementType
	switch elementType {
	case "MET_UCHAR":
		img.pixelType = PixelTypeUInt8
	case "MET_CHAR":
		img.pixelType = PixelTypeInt8
	case "MET_USHORT":
		img.pixelType = PixelTypeUInt16
	case "MET_SHORT":
		img.pixelType = PixelTypeInt16
	case "MET_UINT":
		img.pixelType = PixelTypeUInt32
	case "MET_INT":
		img.pixelType = PixelTypeInt32
	case "MET_ULONG":
		img.pixelType = PixelTypeUInt64
	case "MET_LONG":
		img.pixelType = PixelTypeInt64
	case "MET_FLOAT":
		img.pixelType = PixelTypeFloat32
	case "MET_DOUBLE":
		img.pixelType = PixelTypeFloat64
	default:
		return nil, fmt.Errorf("unsupported element type: %s", elementType)
	}

	// Read the raw data file
	rawPath := filepath.Join(filepath.Dir(filename), rawFilename)
	rawImg, err := readImageTypeRaw(rawPath, img.pixelType)
	if err != nil {
		return nil, fmt.Errorf("failed to read raw data: %v", err)
	}

	// Copy pixel data from raw image
	img.pixels = rawImg.pixels

	return img, nil
}

// WriteImage saves an image to a file.
//
// Parameters:
//   - img: The image object to save
//   - filename: Path to the file where the image will be saved
//   - imageType: Type of image file format
//
// Returns:
//   - error: Error if saving fails
func WriteImage(img *Image, filename string, imageType int) error {
	return img.Save(filename, imageType)
}

// Save saves the image to a file based on the specified image type.
//
// Parameters:
//   - filename: Path to the file where the image will be saved
//   - imageType: Type of image file format
//
// Returns:
//   - error: Error if saving fails
func (img *Image) Save(filename string, imageType int) error {
	switch imageType {
	case ImageTypeRaw:
		return img.saveImageTypeRaw(filename)
	case ImageTypeMHD:
		return img.saveImageTypeMHD(filename)
	default:
		return fmt.Errorf("unknown image type")
	}
}

func (img *Image) saveImageTypeRaw(filename string) error {
	outputFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	switch img.pixelType {
	case PixelTypeUInt8:
		for i := 0; i < len(img.pixels.([]uint8)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]uint8)[i])
		}
	case PixelTypeInt8:
		for i := 0; i < len(img.pixels.([]int8)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]int8)[i])
		}
	case PixelTypeUInt16:
		for i := 0; i < len(img.pixels.([]uint16)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]uint16)[i])
		}
	case PixelTypeInt16:
		for i := 0; i < len(img.pixels.([]int16)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]int16)[i])
		}
	case PixelTypeUInt32:
		for i := 0; i < len(img.pixels.([]uint32)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]uint32)[i])
		}
	case PixelTypeInt32:
		for i := 0; i < len(img.pixels.([]int32)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]int32)[i])
		}
	case PixelTypeUInt64:
		for i := 0; i < len(img.pixels.([]uint64)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]uint64)[i])
		}
	case PixelTypeInt64:
		for i := 0; i < len(img.pixels.([]int64)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]int64)[i])
		}
	case PixelTypeFloat32:
		for i := 0; i < len(img.pixels.([]float32)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]float32)[i])
		}
	case PixelTypeFloat64:
		for i := 0; i < len(img.pixels.([]float64)); i++ {
			binary.Write(outputFile, binary.LittleEndian, img.pixels.([]float64)[i])
		}
	default:
		return fmt.Errorf("unknown pixel type")
	}
	return nil
}

func (img *Image) saveImageTypeMHD(filename string) error {
	// Save the raw data file
	rawFilename := filename[:len(filename)-4] + ".raw"
	if err := img.saveImageTypeRaw(rawFilename); err != nil {
		return fmt.Errorf("failed to save raw data: %v", err)
	}

	// Create and write the MHD header file
	headerFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create MHD file: %v", err)
	}
	defer headerFile.Close()

	// Write header information
	fmt.Fprintf(headerFile, "ObjectType = Image\n")
	fmt.Fprintf(headerFile, "BinaryData = True\n")
	fmt.Fprintf(headerFile, "NDims = %d\n", img.dimension)
	fmt.Fprintf(headerFile, "DimSize = %d", img.size[0])
	for i := 1; i < int(img.dimension); i++ {
		fmt.Fprintf(headerFile, " %d", img.size[i])
	}
	fmt.Fprintf(headerFile, "\n")
	fmt.Fprintf(headerFile, "TransformMatrix = ")
	for i := 0; i < 9; i++ {
		fmt.Fprintf(headerFile, "%f ", img.direction[i])
	}
	fmt.Fprintf(headerFile, "\n")
	fmt.Fprintf(headerFile, "Offset = ")
	for i := 0; i < int(img.dimension); i++ {
		fmt.Fprintf(headerFile, "%f ", img.origin[i])
	}
	fmt.Fprintf(headerFile, "\n")

	// Map pixel type to MHD ElementType
	elementType := ""
	switch img.pixelType {
	case PixelTypeUInt8:
		elementType = "MET_UCHAR"
	case PixelTypeInt8:
		elementType = "MET_CHAR"
	case PixelTypeUInt16:
		elementType = "MET_USHORT"
	case PixelTypeInt16:
		elementType = "MET_SHORT"
	case PixelTypeUInt32:
		elementType = "MET_UINT"
	case PixelTypeInt32:
		elementType = "MET_INT"
	case PixelTypeUInt64:
		elementType = "MET_ULONG"
	case PixelTypeInt64:
		elementType = "MET_LONG"
	case PixelTypeFloat32:
		elementType = "MET_FLOAT"
	case PixelTypeFloat64:
		elementType = "MET_DOUBLE"
	default:
		return fmt.Errorf("unsupported pixel type for MHD format")
	}
	fmt.Fprintf(headerFile, "ElementType = %s\n", elementType)

	// Write additional metadata
	fmt.Fprintf(headerFile, "ElementByteOrderMSB = False\n")
	fmt.Fprintf(headerFile, "CompressedData = False\n")
	fmt.Fprintf(headerFile, "ElementSpacing = %f", img.spacing[0])
	for i := 1; i < int(img.dimension); i++ {
		fmt.Fprintf(headerFile, " %f", img.spacing[i])
	}
	fmt.Fprintf(headerFile, "\n")

	// Reference the raw data file
	rawFilenameBase := rawFilename[len(filepath.Dir(rawFilename))+1:]
	fmt.Fprintf(headerFile, "ElementDataFile = %s\n", rawFilenameBase)

	return nil
}
