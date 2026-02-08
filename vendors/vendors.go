package vendors

import "github.com/startracex/ciinfo/syntax"

type Vendor struct {
	Name     string         `json:"name"`
	Constant string         `json:"constant"`
	Env      syntax.EnvList `json:"env"`
	PR       *syntax.PR     `json:"pr"`
}
