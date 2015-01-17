package implementations

import (
	"math/big"
)

type builtinType struct{}

var Builtin = builtinType{}

func (p builtinType) Pascal(n, m int) int {
	z := big.NewInt(0)
	z.Binomial(int64(n), int64(m))
	return int(z.Int64())
}

func (p builtinType) String() string {
	return "Builtin Implementation"
}
