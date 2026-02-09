package syntax

import (
	"encoding/json"
	"testing"
)

func TestEnvUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
		want Env
	}{
		{
			name: "string",
			data: `"FOO"`,
			want: Env{StrictEqual: "FOO"},
		},
		{
			name: "object with env",
			data: `{"env":"BAR"}`,
			want: Env{StrictEqual: "BAR"},
		},
		{
			name: "object with env and includes",
			data: `{"env":"PATH","includes":"/bin"}`,
			want: Env{StrictEqual: "PATH", Includes: "/bin"},
		},
		{
			name: "object with any",
			data: `{"any":["FOO","BAR"]}`,
			want: Env{EqualsAnyOf: []string{"FOO", "BAR"}},
		},
		{
			name: "object with extra key/value",
			data: `{"FOO":"1","BAR":"2"}`,
			want: Env{EqualsMap: map[string]string{"FOO": "1", "BAR": "2"}},
		},
		{
			name: "mixed",
			data: `{"env":"FOO","includes":"bar","BAZ":"val","any":["X","Y"]}`,
			want: Env{
				StrictEqual: "FOO",
				Includes:    "bar",
				EqualsAnyOf: []string{"X", "Y"},
				EqualsMap:   map[string]string{"BAZ": "val"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e Env
			if err := json.Unmarshal([]byte(tt.data), &e); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			// compare fields individually
			if e.StrictEqual != tt.want.StrictEqual {
				t.Errorf("StrictEqual = %q, want %q", e.StrictEqual, tt.want.StrictEqual)
			}
			if e.Includes != tt.want.Includes {
				t.Errorf("Includes = %q, want %q", e.Includes, tt.want.Includes)
			}
			if len(e.EqualsAnyOf) != len(tt.want.EqualsAnyOf) {
				t.Errorf("EqualsAnyOf = %v, want %v", e.EqualsAnyOf, tt.want.EqualsAnyOf)
			}
			for i := range e.EqualsAnyOf {
				if e.EqualsAnyOf[i] != tt.want.EqualsAnyOf[i] {
					t.Errorf("EqualsAnyOf[%d] = %q, want %q", i, e.EqualsAnyOf[i], tt.want.EqualsAnyOf[i])
				}
			}
			if len(e.EqualsMap) != len(tt.want.EqualsMap) {
				t.Errorf("EqualsMap = %v, want %v", e.EqualsMap, tt.want.EqualsMap)
			}
			for k, v := range tt.want.EqualsMap {
				if e.EqualsMap[k] != v {
					t.Errorf("EqualsMap[%q] = %q, want %q", k, e.EqualsMap[k], v)
				}
			}
		})
	}
}

func TestEnvListUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
		want EnvList
	}{
		{
			name: "single object",
			data: `{"env":"FOO"}`,
			want: EnvList{{StrictEqual: "FOO"}},
		},
		{
			name: "array of objects",
			data: `[{"env":"FOO"},{"any":["X","Y"]}]`,
			want: EnvList{
				{StrictEqual: "FOO"},
				{EqualsAnyOf: []string{"X", "Y"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l EnvList
			if err := json.Unmarshal([]byte(tt.data), &l); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if len(l) != len(tt.want) {
				t.Fatalf("length = %d, want %d", len(l), len(tt.want))
			}
		})
	}
}

func TestPRUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
		want PR
	}{
		{
			name: "string",
			data: `"FOO"`,
			want: PR{StrictEqual: "FOO"},
		},
		{
			name: "object with env",
			data: `{"env":"BAR"}`,
			want: PR{StrictEqual: "BAR"},
		},
		{
			name: "object with not equal",
			data: `{"env":"FOO","ne":"0"}`,
			want: PR{StrictEqual: "FOO", NotEqual: "0"},
		},
		{
			name: "object with any",
			data: `{"any":["FOO","BAR"]}`,
			want: PR{EqualsAnyOf: []string{"FOO", "BAR"}},
		},
		{
			name: "object with extra key/value",
			data: `{"FOO":"1","BAR":"2"}`,
			want: PR{EqualsMap: map[string]string{"FOO": "1", "BAR": "2"}},
		},
		{
			name: "mixed",
			data: `{"env":"FOO","ne":"0","BAZ":"val","any":["X","Y"]}`,
			want: PR{
				StrictEqual: "FOO",
				NotEqual:    "0",
				EqualsAnyOf: []string{"X", "Y"},
				EqualsMap:   map[string]string{"BAZ": "val"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p PR
			if err := json.Unmarshal([]byte(tt.data), &p); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if p.StrictEqual != tt.want.StrictEqual {
				t.Errorf("StrictEqual = %q, want %q", p.StrictEqual, tt.want.StrictEqual)
			}
			if p.NotEqual != tt.want.NotEqual {
				t.Errorf("NotEqual = %q, want %q", p.NotEqual, tt.want.NotEqual)
			}
			if len(p.EqualsAnyOf) != len(tt.want.EqualsAnyOf) {
				t.Errorf("EqualsAnyOf = %v, want %v", p.EqualsAnyOf, tt.want.EqualsAnyOf)
			}
			for i := range p.EqualsAnyOf {
				if p.EqualsAnyOf[i] != tt.want.EqualsAnyOf[i] {
					t.Errorf("EqualsAnyOf[%d] = %q, want %q", i, p.EqualsAnyOf[i], tt.want.EqualsAnyOf[i])
				}
			}
			if len(p.EqualsMap) != len(tt.want.EqualsMap) {
				t.Errorf("EqualsMap = %v, want %v", p.EqualsMap, tt.want.EqualsMap)
			}
			for k, v := range tt.want.EqualsMap {
				if p.EqualsMap[k] != v {
					t.Errorf("EqualsMap[%q] = %q, want %q", k, p.EqualsMap[k], v)
				}
			}
		})
	}
}
