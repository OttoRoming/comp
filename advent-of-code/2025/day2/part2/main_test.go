package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 11+22, solve("11-22"))
	assert.Equal(t, 99+111, solve("95-115"))
	assert.Equal(t, 999+1010, solve("998-1012"))
	assert.Equal(t, 1188511885, solve("1188511880-1188511890"))
	assert.Equal(t, 222222, solve("222220-222224"))
	assert.Equal(t, 0, solve("1698522-1698528"))
	assert.Equal(t, 446446, solve("446443-446449"))
	assert.Equal(t, 38593859, solve("38593856-38593862"))
	assert.Equal(t, 565656, solve("565653-565659"))
	assert.Equal(t, 824824824, solve("824824821-824824827"))
	assert.Equal(t, 2121212121, solve("2121212118-2121212124"))

	assert.Equal(t, 4174379265, solve("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"))
}
