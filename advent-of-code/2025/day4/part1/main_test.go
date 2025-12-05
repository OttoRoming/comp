package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 13, solve("..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."))
}
