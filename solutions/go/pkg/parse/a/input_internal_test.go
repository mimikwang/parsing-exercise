package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputValidate(t *testing.T) {
	cases := []struct {
		name      string
		input     Input
		expectErr bool
	}{
		{
			"Should pass for a valid input",
			Input{
				Id:      1,
				Name:    "name",
				Country: "USA",
				Age:     23,
				Weight:  10.1,
				Height:  23.1,
			},
			false,
		},
		{
			"Should fail for an input with missing values",
			Input{
				Id:     2,
				Age:    23,
				Weight: 21,
			},
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.input.validate()
			if c.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
