package imagetk

import (
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
