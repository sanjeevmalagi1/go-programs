package main_test

import (
	"testing"

	gen "github.com/sanjeevmalagi1/go_programs/15_generics"
)

func TestPrintFormat(t *testing.T) {
	t.Run("Checks format", func(t *testing.T) {
		res := gen.PrintFormat(12)
		if res != "int" {
			t.Errorf("expected format int; got %v", res)
		}
	})

}
