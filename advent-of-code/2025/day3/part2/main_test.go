package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 987654321111, solve("987654321111111"))
	assert.Equal(t, 811111111119, solve("811111111111119"))
	assert.Equal(t, 434234234278, solve("234234234234278"))
	assert.Equal(t, 888911112111, solve("818181911112111"))

	assert.Equal(t, 3121910778619, solve("987654321111111\n811111111111119\n234234234234278\n818181911112111"))
}
