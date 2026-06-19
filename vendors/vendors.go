package vendors

import "github.com/startracex/ciinfo/syntax"

type Vendor struct {
	PR       *syntax.PR     `json:"pr"`
	Env      syntax.EnvList `json:"env"`
	Constant string         `json:"constant"`
	Name     string         `json:"name"`
}
