package common_test

import (
	"testing"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/common"
	"github.com/stretchr/testify/assert"
)

func TestBmi(t *testing.T) {
	cases := []struct {
		name     string
		bmi      float64
		expected string
	}{
		{
			"Should return underweight",
			5.3,
			"underweight",
		},
		{
			"Should return normal",
			20.2,
			"normal",
		},
		{
			"Should return overweight",
			28.8,
			"overweight",
		},
		{
			"Should return very overweight",
			100.2,
			"very overweight",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := common.Bmi(c.bmi)
			assert.Equal(t, c.expected, actual)
		})
	}
}
