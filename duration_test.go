package utils_test

import (
	"testing"

	"github.com/ideatocode/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestDurationUnmarshalText(t *testing.T) {

	x := "61s"

	d := &utils.Duration{}

	d.UnmarshalText([]byte(x))

	assert.Equal(t, float64(61), d.Seconds())
}
