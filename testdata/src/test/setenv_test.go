package test

import (
	"testing"
)

func Test_SetEnv_Func1(t *testing.T) {
	teardown := setup("Test_Setenv_Func1")
	t.Cleanup(teardown)

	t.Setenv("foo", "bar")

	t.Run("Setenv_Func1_Sub1", func(t *testing.T) {
		call("Setenv_Func1_Sub1")
		t.Parallel()
	})

	t.Run("Setenv_Func1_Sub2", func(t *testing.T) {
		call("Setenv_Func1_Sub2")
		t.Parallel()
	})
}
