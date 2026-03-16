package utils_test

import (
	"testing"

	utils "github.com/helios-live/go-utils/v2"
	"github.com/stretchr/testify/assert"
)

func TestDurationUnmarshalText(t *testing.T) {

	x := "61s"

	d := &utils.Duration{}

	d.UnmarshalText([]byte(x))

	assert.Equal(t, float64(61), d.Seconds())
}
