package mypoint_test

import (
	"github.com/stretchr/testify/assert"
	"store-service/internal/mypoint"
	"testing"
)

func Test_sum_1_and_2_equal_3(t *testing.T) {
	sum, error := mypoint.Calculate_num(1, 2)
	if error != nil {
		assert.Error(t, error)
	} else {
		assert.Nil(t, error)
		assert.Equal(t, 3, sum)
	}

}

func Test_sum_minus_1_and_2_error(t *testing.T) {
	sum, error := mypoint.Calculate_num(-1, 2)
	if error != nil {
		assert.Error(t, error, "positive int only, now a = -1")
	} else {
		assert.Nil(t, error)
		assert.Equal(t, 3, sum)
	}

}
