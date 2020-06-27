package utils_test

import (
	"TSACodingChallengeAPI/src/utils"
	"fmt"
	"testing"
)

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{

			"0488445688", "0488445688",
		},
		{
			"+61488224568", "+61488224568",
		},
		{
			"(03) 9333 7119", "0393337119",
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := utils.NormalizePhoneNumber(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestNameToTitleCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{

			"test name", "Test Name",
		},
		{
			"TEST NaMe", "Test Name",
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := utils.NameToTitleCase(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
