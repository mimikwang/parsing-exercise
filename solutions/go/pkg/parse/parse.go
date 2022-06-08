package parse

import (
	"github.com/mimikwang/parsing-exercise/solutions/go/pkg/parse/output"
)

// Parse returns a parsed output
type Parse interface {
	Parse() (output.Output, error)
}
