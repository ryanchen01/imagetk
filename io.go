package imagetk

import (
	"encoding/binary"
	"fmt"
	"os"
)

func (img *Image) Save(filename string) error {
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
