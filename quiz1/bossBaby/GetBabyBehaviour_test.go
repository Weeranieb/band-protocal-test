package bossbaby_test

import (
	bossbaby "bandProtocol/quiz1/bossBaby"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBabyBehaviour(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		expected string
	}{
		{"Action starts with S and end with R", "SR", bossbaby.Good},
		{"Action starts with R and end with S", "RS", bossbaby.Bad},
		{"Empty action", "", bossbaby.Bad},
		{"Single action", "S", bossbaby.Bad},

		{"Revenge more than Shots with extra shot", "SSSRRRRR", bossbaby.Good},
		{"Revenge less than Shots", "SSSRR", bossbaby.Bad},
		{"Revenge more than Shots at first but get shot again", "SSSRRRRSR", bossbaby.Bad},

		{"Shot less but revenge with extra shot", "SSSRSRRRR", bossbaby.Good},
		{"Shot less but revenge with no extra shot", "SSSRSRRR", bossbaby.Good},
		{"Swap and Shot less but revenge less than shot", "SSSRSRR", bossbaby.Bad},

		{"Test case from the quiz 1", "SRSSRRR", bossbaby.Good},
		{"Test case from the quiz 2", "RSSRR", bossbaby.Bad},
		{"Test case from the quiz 3", "SSSRRRRS", bossbaby.Bad},
		{"Test case from the quiz 4", "SRRSSR", bossbaby.Bad},
		{"Test case from the quiz 5", "SSRSRR", bossbaby.Good},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bossbaby.GetBabyBehaviour(tt.action)
			assert.Equal(t, tt.expected, result)
		})
	}
}
