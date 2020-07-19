package options

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		Init("not possible!")
	}, "Init should panic on invalid key")

	assert.NotPanics(func() {
		Init()
	}, "Empty Init should at least initialize logging")
	os.Unsetenv("LOGGING_LEVEL")
	Init()
	assert.Equal(logrus.InfoLevel, logrus.GetLevel(), "Default log level should be info")
}

func TestHas(t *testing.T) {
	assert := assert.New(t)
	a := []string{"a", "b"}
	assert.False(has(a, "z"))
	assert.True(has(a, "a"))
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	a := []string{"a", "b"}
	expected := []string{"c", "a", "b"}
	assert.EqualValues(expected, insert(a, "c"))
}
