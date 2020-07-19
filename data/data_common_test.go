package data

import (
	"testing"

	"github.com/PerforMance308/test2/util"
	"github.com/stretchr/testify/assert"
)

func TestGetCommonEntry(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/common.json", GetCommonEntrys())
	assert.NoError(err)

	assert.Equal(GetCommonEntry("rate"), Common{"rate", 50})
}

func TestGetCommonEntrys(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/common.json", GetCommonEntrys())
	assert.NoError(err)

	assert.Equal(len(GetCommonEntrys().Entrys), 1)
}
