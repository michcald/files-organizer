package internal

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		items         []int
		expectedCount int
	}{
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{},
			0,
		},
	}

	for i, tt := range tests {
		cnt := Count(tt.items...)

		if cnt != tt.expectedCount {
			t.Errorf("test %d: expected count %d. got=%d", i, tt.expectedCount, cnt)
		}
	}
}
