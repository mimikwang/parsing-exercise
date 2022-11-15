package b

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
			"Should not return an error for a valid input",
			Input{
				Id: 23,
				Info: Info{
					Name:    "Pete",
					Country: "Canada",
				},
				Data: Data{
					Bmi: 43.1,
					Age: 20,
				},
			},
			false,
		},
		{
			"Should return an error for input with missing fields",
			Input{
				Id: 42,
				Info: Info{
					Name: "Rita",
				},
				Data: Data{
					Age: 40,
				},
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
