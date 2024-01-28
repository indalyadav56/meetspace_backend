package tests

import (
	"testing"

	"github.com/go-playground/assert/v2"
)


func TestFoo(t *testing.T) {
	assert.Equal(t, "Foo", "Foo")
}

func TestFoo22(t *testing.T) {
	assert.Equal(t, "Foo", "Fo2o")
}
