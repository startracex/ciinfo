package syntax

import (
	"testing"
)

func TestEnvMatch(t *testing.T) {
	tests := []struct {
		name string
		env  Env
		data map[string]string
		want bool
	}{
		{
			name: "StrictEqual exists",
			env:  Env{StrictEqual: "FOO"},
			data: map[string]string{"FOO": "bar"},
			want: true,
		},
		{
			name: "StrictEqual missing",
			env:  Env{StrictEqual: "FOO"},
			data: map[string]string{"BAR": "baz"},
			want: false,
		},
		{
			name: "Includes substring match",
			env:  Env{StrictEqual: "PATH", Includes: "/bin"},
			data: map[string]string{"PATH": "/usr/bin:/usr/local/bin"},
			want: true,
		},
		{
			name: "Includes no match",
			env:  Env{StrictEqual: "PATH", Includes: "/sbin"},
			data: map[string]string{"PATH": "/usr/bin:/usr/local/bin"},
			want: false,
		},
		{
			name: "EqualsAnyOf match",
			env:  Env{EqualsAnyOf: []string{"FOO", "BAR"}},
			data: map[string]string{"BAR": "yes"},
			want: true,
		},
		{
			name: "EqualsAnyOf no match",
			env:  Env{EqualsAnyOf: []string{"FOO", "BAR"}},
			data: map[string]string{"BAZ": "no"},
			want: false,
		},
		{
			name: "EqualsMap match",
			env:  Env{EqualsMap: map[string]string{"FOO": "1", "BAR": "2"}},
			data: map[string]string{"FOO": "1", "BAR": "2"},
			want: true,
		},
		{
			name: "EqualsMap mismatch",
			env:  Env{EqualsMap: map[string]string{"FOO": "1", "BAR": "2"}},
			data: map[string]string{"FOO": "1", "BAR": "3"},
			want: false,
		},
		{
			name: "Empty Env returns false",
			env:  Env{},
			data: map[string]string{"FOO": "1"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.env.Match(tt.data)
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvListMatch(t *testing.T) {
	list := EnvList{
		{StrictEqual: "FOO"},
		{EqualsMap: map[string]string{"BAR": "1"}},
	}

	tests := []struct {
		name string
		data map[string]string
		want bool
	}{
		{
			name: "All match",
			data: map[string]string{"FOO": "x", "BAR": "1"},
			want: true,
		},
		{
			name: "One mismatch",
			data: map[string]string{"FOO": "x", "BAR": "2"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.Match(tt.data)
			if got != tt.want {
				t.Errorf("EnvList.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPRMatch(t *testing.T) {
	tests := []struct {
		name string
		pr   PR
		data map[string]string
		want bool
	}{
		{
			name: "StrictEqual exists",
			pr:   PR{StrictEqual: "FOO"},
			data: map[string]string{"FOO": "1"},
			want: true,
		},
		{
			name: "StrictEqual missing",
			pr:   PR{StrictEqual: "FOO"},
			data: map[string]string{"BAR": "1"},
			want: false,
		},
		{
			name: "StrictEqual + NotEqual match",
			pr:   PR{StrictEqual: "FOO", NotEqual: "0"},
			data: map[string]string{"FOO": "1"},
			want: true,
		},
		{
			name: "StrictEqual + NotEqual mismatch",
			pr:   PR{StrictEqual: "FOO", NotEqual: "0"},
			data: map[string]string{"FOO": "0"},
			want: false,
		},
		{
			name: "EqualsAnyOf match",
			pr:   PR{EqualsAnyOf: []string{"FOO", "BAR"}},
			data: map[string]string{"BAR": "1"},
			want: true,
		},
		{
			name: "EqualsAnyOf no match",
			pr:   PR{EqualsAnyOf: []string{"FOO", "BAR"}},
			data: map[string]string{"BAZ": "1"},
			want: false,
		},
		{
			name: "EqualsMap match",
			pr:   PR{EqualsMap: map[string]string{"FOO": "x", "BAR": "y"}},
			data: map[string]string{"FOO": "x", "BAR": "y"},
			want: true,
		},
		{
			name: "EqualsMap mismatch",
			pr:   PR{EqualsMap: map[string]string{"FOO": "x", "BAR": "y"}},
			data: map[string]string{"FOO": "x", "BAR": "z"},
			want: false,
		},
		{
			name: "Empty PR returns false",
			pr:   PR{},
			data: map[string]string{"FOO": "1"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pr.Match(tt.data)
			if got != tt.want {
				t.Errorf("PR.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
