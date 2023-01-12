package lib

import "testing"

func TestPow(t *testing.T) {
	tt := []struct {
		name string
		x    int
		y    int
		v    int
	}{
		{"two power three", 2, 3, 8},
		{"two power four", 2, 4, 16},
		{"two power one", 2, 1, 2},
	}

	for _, ti := range tt {
		t.Run(ti.name, func(t *testing.T) {
			v := Pow(ti.x, ti.y)

			if v != ti.v {
				t.Errorf("Expected %d and Got %d", ti.v, v)
			}
		})

	}

}
