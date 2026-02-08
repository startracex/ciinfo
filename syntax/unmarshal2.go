//go:build goexperiment.jsonv2
// +build goexperiment.jsonv2

package syntax

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
	ErrInvalidPR  = errors.New("invalid pr")
)

func (e *Env) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {

	case '"': // string
		return json.UnmarshalDecode(dec, &e.StrictEqual)

	case '{': // object
		if _, err := dec.ReadToken(); err != nil {
			return err
		}

		var eq map[string]string

		for dec.PeekKind() != '}' {
			var key string
			if err := json.UnmarshalDecode(dec, &key); err != nil {
				return err
			}

			switch key {
			case "env":
				if err := json.UnmarshalDecode(dec, &e.StrictEqual); err != nil {
					return err
				}
			case "includes":
				if err := json.UnmarshalDecode(dec, &e.Includes); err != nil {
					return err
				}
			case "any":
				if err := json.UnmarshalDecode(dec, &e.EqualsAnyOf); err != nil {
					return err
				}
			default:
				if eq == nil {
					eq = make(map[string]string)
				}
				var val string
				if err := json.UnmarshalDecode(dec, &val); err != nil {
					return err
				}
				eq[key] = val
			}
		}

		if _, err := dec.ReadToken(); err != nil {
			return err
		}

		e.EqualsMap = eq
		return nil
	}

	return ErrInvalidEnv
}

func (l *EnvList) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {

	case '[':
		if _, err := dec.ReadToken(); err != nil { // [
			return err
		}

		var list []Env
		for dec.PeekKind() != ']' {
			var e Env
			if err := e.UnmarshalJSONFrom(dec); err != nil {
				return err
			}
			list = append(list, e)
		}

		if _, err := dec.ReadToken(); err != nil { // ]
			return err
		}

		*l = list
		return nil

	default:
		var single Env
		if err := single.UnmarshalJSONFrom(dec); err != nil {
			return err
		}
		*l = []Env{single}
		return nil
	}
}

func (p *PR) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {

	case '"': // string
		return json.UnmarshalDecode(dec, &p.StrictEqual)

	case '{':
		if _, err := dec.ReadToken(); err != nil { // {
			return err
		}

		var eq map[string]string

		for dec.PeekKind() != '}' {
			var key string
			if err := json.UnmarshalDecode(dec, &key); err != nil {
				return err
			}

			switch key {
			case "env":
				if err := json.UnmarshalDecode(dec, &p.StrictEqual); err != nil {
					return err
				}
			case "any":
				if err := json.UnmarshalDecode(dec, &p.EqualsAnyOf); err != nil {
					return err
				}
			case "ne":
				if err := json.UnmarshalDecode(dec, &p.NotEqual); err != nil {
					return err
				}
			default:
				if eq == nil {
					eq = make(map[string]string)
				}
				var val string
				if err := json.UnmarshalDecode(dec, &val); err != nil {
					return err
				}
				eq[key] = val
			}
		}

		if _, err := dec.ReadToken(); err != nil { // }
			return err
		}

		p.EqualsMap = eq
		return nil
	}

	return ErrInvalidPR
}
