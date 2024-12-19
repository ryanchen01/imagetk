package imagetk

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestInvert2x2(t *testing.T) {
	tests := []struct {
		name    string
		input   [4]float64
		want    [4]float64
		wantErr bool
	}{
		{
			name:    "identity matrix",
			input:   [4]float64{1, 0, 0, 1},
			want:    [4]float64{1, 0, 0, 1},
			wantErr: false,
		},
		{
			name:    "regular matrix",
			input:   [4]float64{4, 7, 2, 6},
			want:    [4]float64{0.6, -0.7, -0.2, 0.4},
			wantErr: false,
		},
		{
			name:    "singular matrix",
			input:   [4]float64{1, 2, 2, 4},
			want:    [4]float64{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := invert2x2(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("invert2x2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check inverse matrix values
				for i := 0; i < 4; i++ {
					if math.Abs(got[i]-tt.want[i]) > 1e-6 {
						t.Errorf("invert2x2() = %v, want %v", got, tt.want)
						break
					}
				}

				// Verify multiplication with original matrix gives identity
				result := [4]float64{
					got[0]*tt.input[0] + got[1]*tt.input[2],
					got[0]*tt.input[1] + got[1]*tt.input[3],
					got[2]*tt.input[0] + got[3]*tt.input[2],
					got[2]*tt.input[1] + got[3]*tt.input[3],
				}
				identity := [4]float64{1, 0, 0, 1}
				for i := 0; i < 4; i++ {
					if math.Abs(result[i]-identity[i]) > 1e-6 {
						t.Errorf("Matrix multiplication with inverse failed, got %v, want identity", result)
						break
					}
				}
			}
		})
	}
}

func TestInvert3x3(t *testing.T) {
	tests := []struct {
		name    string
		input   [9]float64
		want    [9]float64
		wantErr bool
	}{
		{
			name:    "identity matrix",
			input:   [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			want:    [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			wantErr: false,
		},
		{
			name:    "regular matrix",
			input:   [9]float64{1, 2, 3, 0, 1, 4, 5, 6, 0},
			want:    [9]float64{-24, 18, 5, 20, -15, -4, -5, 4, 1},
			wantErr: false,
		},
		{
			name:    "singular matrix",
			input:   [9]float64{1, 2, 3, 2, 4, 6, 3, 6, 9},
			want:    [9]float64{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := invert3x3(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("invert3x3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check inverse matrix values
				for i := 0; i < 9; i++ {
					if math.Abs(got[i]-tt.want[i]) > 1e-6 {
						t.Errorf("invert3x3() = %v, want %v", got, tt.want)
						break
					}
				}

				// Verify multiplication with original matrix gives identity
				result := [9]float64{
					got[0]*tt.input[0] + got[1]*tt.input[3] + got[2]*tt.input[6],
					got[0]*tt.input[1] + got[1]*tt.input[4] + got[2]*tt.input[7],
					got[0]*tt.input[2] + got[1]*tt.input[5] + got[2]*tt.input[8],
					got[3]*tt.input[0] + got[4]*tt.input[3] + got[5]*tt.input[6],
					got[3]*tt.input[1] + got[4]*tt.input[4] + got[5]*tt.input[7],
					got[3]*tt.input[2] + got[4]*tt.input[5] + got[5]*tt.input[8],
					got[6]*tt.input[0] + got[7]*tt.input[3] + got[8]*tt.input[6],
					got[6]*tt.input[1] + got[7]*tt.input[4] + got[8]*tt.input[7],
					got[6]*tt.input[2] + got[7]*tt.input[5] + got[8]*tt.input[8],
				}
				identity := [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
				for i := 0; i < 9; i++ {
					if math.Abs(result[i]-identity[i]) > 1e-6 {
						t.Errorf("Matrix multiplication with inverse failed, got %v, want identity", result)
						break
					}
				}
			}
		})
	}
}

func TestGetPixelFromPoint(t *testing.T) {
	// Create a 2D test image with known values
	create2DImage := func(pixelType int) *Image {
		img := &Image{
			dimension: 2,
			size:      []uint32{3, 3},
			spacing:   []float64{1.0, 1.0},
			origin:    []float64{0.0, 0.0},
			direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			pixelType: pixelType,
		}

		// Create test data: [1 2 3; 4 5 6; 7 8 9]
		switch pixelType {
		case PixelTypeUInt8:
			img.pixels = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
		case PixelTypeFloat32:
			data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}
			img.pixels = make([]byte, len(data)*4)
			for i, v := range data {
				binary.LittleEndian.PutUint32(img.pixels[i*4:], math.Float32bits(v))
			}
		}
		return img
	}

	tests := []struct {
		name     string
		img      *Image
		point    []float64
		fillType int
		want     float64
		wantErr  bool
	}{
		{
			name:     "center point uint8",
			img:      create2DImage(PixelTypeUInt8),
			point:    []float64{1.0, 1.0},
			fillType: FillTypeNearest,
			want:     5.0,
			wantErr:  false,
		},
		{
			name:     "interpolated point uint8",
			img:      create2DImage(PixelTypeUInt8),
			point:    []float64{0.5, 0.5},
			fillType: FillTypeNearest,
			want:     3.0, // Average of 1,2,4,5
			wantErr:  false,
		},
		{
			name:     "outside point with zero fill",
			img:      create2DImage(PixelTypeUInt8),
			point:    []float64{-1.0, -1.0},
			fillType: FillTypeZero,
			want:     0.0,
			wantErr:  false,
		},
		{
			name:     "outside point with nearest fill",
			img:      create2DImage(PixelTypeUInt8),
			point:    []float64{-0.5, -0.5},
			fillType: FillTypeNearest,
			want:     1.0,
			wantErr:  false,
		},
		{
			name:     "center point float32",
			img:      create2DImage(PixelTypeFloat32),
			point:    []float64{1.0, 1.0},
			fillType: FillTypeNearest,
			want:     5.0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.img.GetPixelFromPoint(tt.point, tt.fillType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPixelFromPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("GetPixelFromPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
