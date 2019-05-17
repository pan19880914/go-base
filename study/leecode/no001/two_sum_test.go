package no001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkCacheData(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b.ResetTimer()
	twoSum(nums, 29)
	b.StopTimer()
}

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	expected := []int{13, 14}
	actual := twoSum(nums, 29)
	assert.Equal(t, expected, actual)
}
