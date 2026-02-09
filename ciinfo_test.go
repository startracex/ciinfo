package ciinfo

import (
	"reflect"
	"testing"

	"github.com/startracex/ciinfo/syntax"
	"github.com/startracex/ciinfo/vendors"
)

func TestGetInfoFrom_CIExplicitFalse(t *testing.T) {
	env := map[string]string{
		"CI": "false",
	}

	info := GetInfoFrom(env, nil)
	if info != nil {
		t.Fatalf("expected nil when CI=false, got %+v", info)
	}
}

func TestGetInfoFrom_VendorMatch(t *testing.T) {
	vlist := []vendors.Vendor{
		{
			Name:     "TestCI",
			Constant: "TEST",
			Env: syntax.EnvList{
				{StrictEqual: "TEST_ENV"},
			},
		},
	}

	env := map[string]string{
		"TEST_ENV": "1",
	}

	info := GetInfoFrom(env, vlist)
	if info == nil {
		t.Fatal("expected info, got nil")
	}

	if !info.IsCI {
		t.Error("IsCI should be true")
	}
	if info.Name != "TestCI" {
		t.Errorf("Name = %q, want TestCI", info.Name)
	}
	if info.ID != "TEST" {
		t.Errorf("ID = %q, want TEST", info.ID)
	}
	if !info.Vendors["TEST"] {
		t.Error("Vendors map not set correctly")
	}
}

func TestGetInfoFrom_VendorNoMatch(t *testing.T) {
	vlist := []vendors.Vendor{
		{
			Name:     "TestCI",
			Constant: "TEST",
			Env: syntax.EnvList{
				{StrictEqual: "TEST_ENV"},
			},
		},
	}

	env := map[string]string{}

	info := GetInfoFrom(env, vlist)
	if info == nil {
		t.Fatal("expected non-nil info (commonKeys may still apply)")
	}

	if info.IsCI {
		t.Error("IsCI should be false when nothing matches")
	}
	if len(info.Vendors) != 0 {
		t.Errorf("Vendors should be empty, got %+v", info.Vendors)
	}
}

func TestGetInfoFrom_PRMatch(t *testing.T) {
	vlist := []vendors.Vendor{
		{
			Name:     "TestCI",
			Constant: "TEST",
			Env: syntax.EnvList{
				{StrictEqual: "CI_VENDOR"},
			},
			PR: &syntax.PR{
				StrictEqual: "PR_FLAG",
			},
		},
	}

	env := map[string]string{
		"CI_VENDOR": "1",
		"PR_FLAG":   "true",
	}

	info := GetInfoFrom(env, vlist)
	if info == nil {
		t.Fatal("expected info, got nil")
	}

	if !info.IsPR {
		t.Error("IsPR should be true")
	}
}

func TestGetInfoFrom_CommonKeysFallback(t *testing.T) {
	env := map[string]string{
		"BUILD_ID": "123",
	}

	info := GetInfoFrom(env, nil)
	if info == nil {
		t.Fatal("expected info, got nil")
	}

	if !info.IsCI {
		t.Error("IsCI should be true from common keys")
	}
	if info.Name != "" || info.ID != "" {
		t.Error("Name and ID should be empty when no vendor matched")
	}
}

func TestGetInfoFrom_MultipleVendorsLastWins(t *testing.T) {
	vlist := []vendors.Vendor{
		{
			Name:     "FirstCI",
			Constant: "FIRST",
			Env: syntax.EnvList{
				{StrictEqual: "A"},
			},
		},
		{
			Name:     "SecondCI",
			Constant: "SECOND",
			Env: syntax.EnvList{
				{StrictEqual: "B"},
			},
		},
	}

	env := map[string]string{
		"A": "1",
		"B": "1",
	}

	info := GetInfoFrom(env, vlist)
	if info == nil {
		t.Fatal("expected info, got nil")
	}

	if info.Name != "SecondCI" || info.ID != "SECOND" {
		t.Errorf("last matching vendor should win, got Name=%q ID=%q", info.Name, info.ID)
	}

	expectedVendors := map[string]bool{
		"FIRST":  true,
		"SECOND": true,
	}
	if !reflect.DeepEqual(info.Vendors, expectedVendors) {
		t.Errorf("Vendors map = %+v, want %+v", info.Vendors, expectedVendors)
	}
}
