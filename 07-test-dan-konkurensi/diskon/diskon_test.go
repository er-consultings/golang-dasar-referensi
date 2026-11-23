package diskon

import "testing"

func TestHitung(t *testing.T) {
	tests := []struct {
		name   string
		total  int64
		persen int64
		want   int64
	}{
		{"tanpa diskon", 100000, 0, 100000},
		{"sepuluh persen", 100000, 10, 90000},
		{"persen negatif diabaikan", 50000, -5, 50000},
		{"lebih dari 100 dibatasi", 50000, 150, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hitung(tt.total, tt.persen); got != tt.want {
				t.Errorf("Hitung(%d, %d) = %d, ingin %d", tt.total, tt.persen, got, tt.want)
			}
		})
	}
}
