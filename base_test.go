package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/vbatts/source-manifest/types"
)

func TestStructHost(t *testing.T) {
	h := types.Host{}
	spew.Dump(h)

	buf, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	t.Fatalf("%q struct", types.StructTypeHOST)
}
func TestStructMaterials(t *testing.T) {
	m := types.Materials{}
	m.Packages = append(m.Packages, types.Package{})
	spew.Dump(m)

	buf, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	t.Fatalf("%q struct", types.StructTypeMATERIALS)
}
func TestStructStep(t *testing.T) {
	s := types.Step{}
	spew.Dump(s)

	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	t.Fatalf("%q struct", types.StructTypeSTEP)
}
