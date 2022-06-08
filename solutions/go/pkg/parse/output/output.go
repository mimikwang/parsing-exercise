package output

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// Output defines the output format
type Output struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
	Bmi  string `json:"bmi" validate:"required"`
}

// NewOutput constructs a new Output
func NewOutput(id int, name string, age int, bmi string) (Output, error) {
	output := Output{
		Id:   id,
		Name: name,
		Age:  age,
		Bmi:  bmi,
	}
	if err := output.validate(); err != nil {
		return Output{}, err
	}
	return output, nil
}

// ToBytes convert Output to bytes
func (o Output) ToBytes() []byte {
	data, _ := json.Marshal(o)
	return data
}

// validate Output's fields
func (o *Output) validate() error {
	return validate.Struct(o)
}
