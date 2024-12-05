package imagetk

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteImage(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_find_prjs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	img, err := NewImage([]uint32{10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Fatalf("failed to create image: %v", err)
	}
	tests := []struct {
		name      string
		img       *Image
		filename  string
		imageType int
		expect    error
	}{
		{name: "write raw image", img: img, filename: filepath.Join(tempDir, "test.raw"), imageType: ImageTypeRaw, expect: nil},
		{name: "write mhd image", img: img, filename: filepath.Join(tempDir, "test.mhd"), imageType: ImageTypeMHD, expect: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteImage(tt.img, tt.filename, tt.imageType)
			if err != tt.expect {
				t.Errorf("expected error %v, got %v", tt.expect, err)
			}
		})
	}
}

func TestReadImage(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_find_prjs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	img, err := NewImage([]uint32{10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Fatalf("failed to create image: %v", err)
	}
	err = img.SetPixel([]uint32{0, 0}, float32(1))
	if err != nil {
		t.Fatalf("failed to set pixel: %v", err)
	}
	err = WriteImage(img, filepath.Join(tempDir, "test.mhd"), ImageTypeMHD)
	if err != nil {
		t.Fatalf("failed to write image: %v", err)
	}
	pixelType := PixelTypeFloat32
	tests := []struct {
		name      string
		filename  string
		imageType int
		expect    error
	}{
		{name: "read raw image", filename: filepath.Join(tempDir, "test.raw"), imageType: ImageTypeRaw, expect: nil},
		{name: "read mhd image", filename: filepath.Join(tempDir, "test.mhd"), imageType: ImageTypeMHD, expect: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := ReadImage(tt.filename, tt.imageType, &pixelType)
			if err != tt.expect {
				t.Errorf("expected error %v, got %v", tt.expect, err)
			}
			if tt.imageType == ImageTypeRaw {
				err = img.SetSize([]uint32{10, 10})
				if err != nil {
					t.Fatalf("failed to set size: %v", err)
				}
			}
			pixel, err := img.GetPixel([]uint32{0, 0})
			if err != nil {
				t.Fatalf("failed to get pixel: %v", err)
			}
			if pixel.(float32) != 1 {
				t.Errorf("expected pixel value 1, got %v", pixel)
			}
		})
	}
}
