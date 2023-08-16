package tests

import (
	"testing"

	"github.com/AulaDevs/Utility"
)

func TestOptions(t *testing.T) {
	test := func(value Utility.Option[string]) {
		t.Logf("The value is none? %v", value.IsNone())
		t.Logf("The value is some? %v", value.IsSome())

		if value.IsSome() {
			t.Logf("The value is %s", value.Unwrap())
		}
	}

	test(Utility.None[string]())
	test(Utility.Some("Hello world"))
}
