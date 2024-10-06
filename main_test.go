package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

type dateTestdata struct {
	Input    string
	Expected time.Time
}

var testDates = []dateTestdata{
	{
		Input:    "2022-01-02",
		Expected: time.Date(2022, 1, 2, 23, 59, 58, 0, time.UTC),
	},
	{
		Input:    "02-01-2022",
		Expected: time.Date(2022, 1, 2, 23, 59, 58, 0, time.UTC),
	},
	{
		Input:    "12.06.2022",
		Expected: time.Date(2022, 6, 12, 23, 59, 58, 0, time.UTC),
	},
	{
		Input:    "12.6.2022",
		Expected: time.Date(2022, 6, 12, 23, 59, 58, 0, time.UTC),
	},
	{
		Input:    "12.33.2022",
		Expected: time.Time{},
	},
	{
		Input:    "12-10-2022 17:32:00",
		Expected: time.Date(2022, 10, 12, 17, 32, 0, 0, time.UTC),
	},
	{
		Input:    "12-10-2022T17:32:00",
		Expected: time.Date(2022, 10, 12, 17, 32, 0, 0, time.UTC),
	},
	{
		Input:    "12,06,2022",
		Expected: time.Date(2022, 6, 12, 23, 59, 58, 0, time.UTC),
	},
	{
		Input:    "12.10.2022 17:32:00",
		Expected: time.Date(2022, 10, 12, 17, 32, 0, 0, time.UTC),
	},
	{
		Input:    "",
		Expected: time.Now(),
	},
	{
		Input:    "abc",
		Expected: time.Time{},
	},
}

func TestStringDateAdapter(t *testing.T) {
	for _, tc := range testDates {
		t.Run(tc.Input, func(t *testing.T) {
			answer := stringDateAdapter(tc.Input)

			require.WithinDuration(t, tc.Expected, answer, 2*time.Second)

		})
	}
}

func TestNewConfiguration(t *testing.T) {
	var defaultVal configurations
	t.Run("newConfiguration", func(t *testing.T) {
		result := newConfiguration()
		if result == defaultVal {
			t.Error("new configuration should return default configuration")
		}

	})
}
func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)

}
