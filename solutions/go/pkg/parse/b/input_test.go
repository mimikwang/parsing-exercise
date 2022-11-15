package b_test

import (
	"testing"

	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/b"
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
				`{"id": 42, "info": {"name": "foo", "country": "mexico"}, "data": {"bmi": 42.1, "age": 21}}`,
			),
			b.Input{
				Id: 42,
				Info: b.Info{
					Name:    "foo",
					Country: "mexico",
				},
				Data: b.Data{
					Bmi: 42.1,
					Age: 21,
				},
			},
			false,
		},
		{
			"Should succeed with a valid input with an extra field",
			[]byte(
				`{"extra": "bla", "id": 42, "info": {"name": "foo", "country": "mexico"}, "data": {"bmi": 42.1, "age": 21}}`,
			),
			b.Input{
				Id: 42,
				Info: b.Info{
					Name:    "foo",
					Country: "mexico",
				},
				Data: b.Data{
					Bmi: 42.1,
					Age: 21,
				},
			},
			false,
		},
		{
			"Should error out for an invalid json",
			[]byte(
				`{hello`,
			),
			b.Input{},
			true,
		},
		{
			"Should error out for missing required fields",
			[]byte(
				`{"id": 42, "info": {"name": "foo"}, "data": {"bmi": 42.1, "age": 21}}`,
			),
			b.Input{},
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := b.NewInput(c.bytes)
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
		input     b.Input
		expected  output.Output
		expectErr bool
	}{
		{
			"Should return the expected parsed output",
			b.Input{
				Id: 42,
				Info: b.Info{
					Name:    "foo",
					Country: "mexico",
				},
				Data: b.Data{
					Bmi: 42.1,
					Age: 21,
				},
			},
			output.Output{
				Id:   42,
				Name: "foo",
				Age:  21,
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
