package greetings_test

import (
	"example_module/src/greetings"
	"testing"
	"github.com/go-playground/assert"
)

func TestGreeting(t *testing.T) {
	assert.NotEqual(t, greetings.Greeting(), "")
}
