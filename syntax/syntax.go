package syntax

import (
	"strings"
)

type EnvList []Env

type Env struct {
	StrictEqual string
	Includes    string
	EqualsAnyOf []string
	EqualsMap   map[string]string
}

type PR struct {
	StrictEqual string
	NotEqual    string
	EqualsAnyOf []string
	EqualsMap   map[string]string
}

func (r *Env) Match(env map[string]string) bool {
	switch {
	case r.StrictEqual != "" && r.Includes == "":
		return env[r.StrictEqual] != ""

	case r.StrictEqual != "" && r.Includes != "":
		return strings.Contains(env[r.StrictEqual], r.Includes)

	case len(r.EqualsAnyOf) > 0:
		for _, k := range r.EqualsAnyOf {
			if env[k] != "" {
				return true
			}
		}
		return false

	case len(r.EqualsMap) > 0:
		for k, v := range r.EqualsMap {
			if env[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

func (l *EnvList) Match(env map[string]string) bool {
	for _, rule := range *l {
		if !rule.Match(env) {
			return false
		}
	}
	return true
}

func (r *PR) Match(env map[string]string) bool {
	switch {
	case r.StrictEqual != "" && len(r.EqualsAnyOf) == 0 && r.NotEqual == "":
		return env[r.StrictEqual] != ""

	case r.StrictEqual != "" && len(r.EqualsAnyOf) > 0:
		for _, v := range r.EqualsAnyOf {
			if env[r.StrictEqual] == v {
				return true
			}
		}
		return false

	case r.StrictEqual != "" && r.NotEqual != "":
		val, ok := env[r.StrictEqual]
		return ok && val != r.NotEqual

	case len(r.EqualsAnyOf) > 0:
		for _, k := range r.EqualsAnyOf {
			if env[k] != "" {
				return true
			}
		}
		return false

	case len(r.EqualsMap) > 0:
		for k, v := range r.EqualsMap {
			if env[k] != v {
				return false
			}
		}
		return true
	}
	return false
}
