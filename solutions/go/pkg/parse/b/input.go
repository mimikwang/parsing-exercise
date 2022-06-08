package b

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/common"
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/output"
)

var (
	validate = validator.New()
)

// Input represents input format B
type Input struct {
	Id   int  `json:"id" validate:"required"`
	Info Info `json:"info" validate:"required"`
	Data Data `json:"data" validate:"required"`
}

// Info represents the info portion of Input
type Info struct {
	Name    string `json:"name" validate:"required"`
	Country string `json:"country" validate:"required"`
}

// Data represents the data portion of Input
type Data struct {
	Bmi float64 `json:"bmi" validate:"required"`
	Age int     `json:"age" validate:"required"`
}

// NewInput constructs a new input from bytes and validates the fields
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
	return output.NewOutput(
		i.Id,
		i.Info.Name,
		i.Data.Age,
		common.Bmi(i.Data.Bmi),
	)
}

// validate Input's fields
func (i *Input) validate() error {
	return validate.Struct(i)
}
