package mypoint_test

import (
	"store-service/internal/mypoint"
	"testing"
	"github.com/stretchr/testify/assert"

)
func Test_100_baht_should_get_1_point(t *testing.T) {
	expected := mypoint.Calculate_point(100)
	actual_point := 1
	assert.Equal(t, expected, actual_point)
}

func Test_99_baht_should_get_0_point(t *testing.T) {
	expected := mypoint.Calculate_point(99)
	actual_point := 0
	assert.Equal(t, expected, actual_point)
}

func Test_299_baht_should_get_2point(t *testing.T) {
	expected := mypoint.Calculate_point(299)
	actual_point := 2
	assert.Equal(t, expected, actual_point)
}