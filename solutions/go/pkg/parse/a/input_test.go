package a_test

import (
	"testing"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/a"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/output"
	"github.com/stretchr/testify/assert"
)

func TestNewInput(t *testing.T) {
	cases := []struct {
		name      string
		bytes     []byte
		expected  parse.Parse
		expectErr bool
	}{
		{
			"Should succeed with a valid input",
			[]byte(
				`{"id": 3, "name": "hello", "country": "USA", "age": 42, "weight": 21.1, "height": 42.4}`,
			),
			a.Input{
				Id:      3,
				Name:    "hello",
				Country: "USA",
				Age:     42,
				Weight:  21.1,
				Height:  42.4,
			},
			false,
		},
		{
			"Should succeed with a valid input with an extra field",
			[]byte(
				`{"id": 3, "name": "hello", "country": "USA", "age": 42, "weight": 21.1, "height": 42.4, "extra": "extra"}`,
			),
			a.Input{
				Id:      3,
				Name:    "hello",
				Country: "USA",
				Age:     42,
				Weight:  21.1,
				Height:  42.4,
			},
			false,
		},
		{
			"Should error out for an invalid json",
			[]byte(
				`{"id": 3`,
			),
			a.Input{},
			true,
		},
		{
			"Should error out for missing required fields",
			[]byte(
				`{"id": 3, "country": "USA", "age": 42, "weight": 21.1, "height": 42.4}`,
			),
			a.Input{},
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := a.NewInput(c.bytes)
			assert.Equal(t, c.expected, actual)
			if c.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestInputParse(t *testing.T) {
	cases := []struct {
		name      string
		input     a.Input
		expected  output.Output
		expectErr bool
	}{
		{
			"Should return the expected parsed output",
			a.Input{
				Id:      21,
				Name:    "foo",
				Country: "bar",
				Age:     42,
				Weight:  210000.1,
				Height:  1,
			},
			output.Output{
				Id:   21,
				Name: "foo",
				Age:  42,
				Bmi:  "very overweight",
			},
			false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.Parse()
			assert.Equal(t, c.expected, actual)
			if c.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
