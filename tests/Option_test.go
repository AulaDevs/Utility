package tests

import (
	"testing"

	"github.com/AulaDevs/Utility/option"
)

func TestOptions(t *testing.T) {
	test := func(value option.Option[string]) {
		t.Logf("Is the value none? %v", value.IsNone())
		t.Logf("Is the value some? %v", value.IsSome())
		t.Logf("Unwrap value: %s", value.UnwrapOr(func() string { return "none" }))
	}

	test(option.None[string]())
	test(option.Some("Hello world"))
}
