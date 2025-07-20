package days

import "testing"

func TestCountDaysFrom(t *testing.T) {
	tests := []struct {
		dateStr string
		wantMin int
	}{
		{"01.01.2020", 1600},
		{"20.07.2020", 1400},
	}

	for _, tt := range tests {
		got, err := CountDaysFrom(tt.dateStr)
		if err != nil {
			t.Errorf("ошибка для %q: %v", tt.dateStr, err)
			continue
		}
		if got < tt.wantMin {
			t.Errorf("для %q: ожидали минимум %d, но получили %d", tt.dateStr, tt.wantMin, got)
		}
	}
}
