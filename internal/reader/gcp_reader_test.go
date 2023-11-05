package reader

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseFrom(t *testing.T) {
	tests := []struct {
		name       string
		givenValue string
		wantsValue string
	}{
		{
			name:       "Test tail",
			givenValue: "tail",
			wantsValue: "tail",
		},
		{
			name:       "Test Relative Second",
			givenValue: "86400s",
			wantsValue: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			name:       "Test Relative Minute",
			givenValue: "1440m",
			wantsValue: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			name:       "Test Relative Hour",
			givenValue: "24h",
			wantsValue: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			name:       "Test Relative Day",
			givenValue: "1d",
			wantsValue: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			name:       "Test Fixed Time",
			givenValue: "2021-01-30T15:00:00",
			wantsValue: func() string {
				tv, _ := time.Parse("2006-01-02T15:04:05", "2021-01-30T15:00:00")
				return tv.Format(time.RFC3339)
			}(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := ParseFrom(test.givenValue)
			fmt.Println(v)
			assert.Equal(t, test.wantsValue, v)
		})
	}
}
