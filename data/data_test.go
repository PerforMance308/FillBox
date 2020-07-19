package data

import (
	"testing"

	"github.com/PerforMance308/test2/util"
	"github.com/stretchr/testify/assert"
)

func TestInitData(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/block.json", GetBlockEntrys())
	assert.NoError(err)

	err = util.ParseJSONFile("../configs/interval.json", GetIntervalEntrys())
	assert.NoError(err)

	err = util.ParseJSONFile("../configs/common.json", GetCommonEntrys())
	assert.NoError(err)

	assert.NotNil(GetBlockEntrys())
	assert.NotNil(GetIntervalEntrys())
	assert.NotNil(GetCommonEntrys())
}
