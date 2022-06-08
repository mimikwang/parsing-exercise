package a

import (
	"encoding/json"
	"math"

	"github.com/go-playground/validator/v10"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/common"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/output"
)

var (
	validate = validator.New()
)

// Input represents input format A
type Input struct {
	Id      int     `json:"id" validate:"required"`
	Name    string  `json:"name" validate:"required"`
	Country string  `json:"country" validate:"required"`
	Age     int     `json:"age" validate:"required"`
	Weight  float64 `json:"weight" validate:"required"`
	Height  float64 `json:"height" validate:"required"`
}

// NewInput constructs a new input from bytes and validate the fields
func NewInput(bytes []byte) (parse.Parse, error) {
	var input Input
	if err := json.Unmarshal(bytes, &input); err != nil {
		return Input{}, err
	}
	if err := input.validate(); err != nil {
		return Input{}, err
	}
	return input, nil
}

// Parse converts Input to an Output
func (i Input) Parse() (output.Output, error) {
	bmi := i.Weight / math.Pow(i.Height, 2.0)
	return output.NewOutput(
		i.Id,
		i.Name,
		i.Age,
		common.Bmi(bmi),
	)
}

// validate Input's fields
func (i *Input) validate() error {
	return validate.Struct(i)
}
