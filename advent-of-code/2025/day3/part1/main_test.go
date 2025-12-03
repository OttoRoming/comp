package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 98, solve("987654321111111"))
	assert.Equal(t, 89, solve("811111111111119"))
	assert.Equal(t, 78, solve("234234234234278"))
	assert.Equal(t, 92, solve("818181911112111"))

	assert.Equal(t, 357, solve("987654321111111\n811111111111119\n234234234234278\n818181911112111"))
}
