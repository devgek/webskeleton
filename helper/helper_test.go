package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kahrersoftware.at/webskeleton/helper"
)

func TestMonthFromDay(t *testing.T) {
	i := helper.MonthFromDay("2019-01-31")
	assert.Equal(t, 1, i, "i must be 1")
}
