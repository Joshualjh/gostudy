package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert.Equal(t, 81, square(9), " test")
}
