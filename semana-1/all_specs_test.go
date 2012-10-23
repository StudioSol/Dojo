package main

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(QuebraLinhaTest)
	gospec.MainGoTest(r, t) }
