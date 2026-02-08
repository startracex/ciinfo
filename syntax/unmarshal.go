//go:build !goexperiment.jsonv2
// +build !goexperiment.jsonv2

package syntax

import "encoding/json"

func (e *Env) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		e.StrictEqual = s
		return nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	e.EqualsMap = make(map[string]string)

	for k, v := range raw {
		switch k {
		case "env":
			if err := json.Unmarshal(v, &e.StrictEqual); err != nil {
				return err
			}
		case "includes":
			if err := json.Unmarshal(v, &e.Includes); err != nil {
				return err
			}
		case "any":
			if err := json.Unmarshal(v, &e.EqualsAnyOf); err != nil {
				return err
			}
		default:
			var val string
			if err := json.Unmarshal(v, &val); err == nil {
				e.EqualsMap[k] = val
			}
		}
	}

	if len(e.EqualsMap) == 0 {
		e.EqualsMap = nil
	}

	return nil
}

func (l *EnvList) UnmarshalJSON(data []byte) error {

	if len(data) > 0 && data[0] == '[' {
		return json.Unmarshal(data, (*[]Env)(l))
	}

	var single Env
	if err := json.Unmarshal(data, &single); err != nil {
		return err
	}
	*l = []Env{single}
	return nil
}

func (p *PR) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		p.StrictEqual = s
		return nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.EqualsMap = make(map[string]string)

	for k, v := range raw {
		switch k {
		case "env":
			if err := json.Unmarshal(v, &p.StrictEqual); err != nil {
				return err
			}
		case "any":
			if err := json.Unmarshal(v, &p.EqualsAnyOf); err != nil {
				return err
			}
		case "ne":
			if err := json.Unmarshal(v, &p.NotEqual); err != nil {
				return err
			}
		default:
			var val string
			if err := json.Unmarshal(v, &val); err == nil {
				p.EqualsMap[k] = val
			}
		}
	}

	if len(p.EqualsMap) == 0 {
		p.EqualsMap = nil
	}

	return nil
}
