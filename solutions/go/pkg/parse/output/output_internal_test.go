package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputValidate(t *testing.T) {
	cases := []struct {
		name      string
		output    Output
		expectErr bool
	}{
		{
			"Should not return an error for a valid output",
			Output{
				Id:   1,
				Name: "name",
				Age:  38,
				Bmi:  "healthy",
			},
			false,
		},
		{
			"Should error for an empty output",
			Output{},
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.output.validate()
			if c.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
