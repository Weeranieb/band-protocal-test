package rescuechicken_test

import (
	"testing"

	rescuechicken "github.com/weeranieb/band-protocal-test/quiz2/rescueChicken"

	"github.com/stretchr/testify/assert"
)

func TestRescueChicken(t *testing.T) {
	tests := []struct {
		chickenNo        int
		roofSize         int
		chickensPosition []int
		expectedDistance int
	}{
		{5, 5, []int{2, 5, 10, 12, 15}, 2},
		{6, 10, []int{1, 11, 30, 34, 35, 37}, 4},

		// edge case
		{0, 5, []int{}, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expectedDistance, rescuechicken.RescueChicken(tt.chickenNo, tt.roofSize, tt.chickensPosition))
	}
}
