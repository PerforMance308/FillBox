package data

import (
	"testing"

	"github.com/PerforMance308/test2/util"
	"github.com/stretchr/testify/assert"
)

func TestGetIntervalEntry(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/interval.json", GetIntervalEntrys())
	assert.NoError(err)

	assert.Equal(GetIntervalEntry(1), Interval{1, []int{1, 2}, false, 10})
	assert.Equal(GetIntervalEntry(7), Interval{7, []int{11, 12, 17, 18}, true, 10})
}

func TestGetIntervalEntrys(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/interval.json", GetIntervalEntrys())
	assert.NoError(err)

	assert.Equal(len(GetIntervalEntrys().Entrys), 7)
}

func TestGetLastInterval(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/interval.json", GetIntervalEntrys())
	assert.NoError(err)

	assert.Equal(GetLastInterval(), Interval{7, []int{11, 12, 17, 18}, true, 10})
}
