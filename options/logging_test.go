package options

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	assert := assert.New(t)

	Init()
	assert.Equal(logrus.InfoLevel, logrus.GetLevel(), "Default log level should be info")

	os.Setenv("LOGGING_LEVEL", "error")
	Init()
	assert.Equal(logrus.ErrorLevel, logrus.GetLevel())
}
