package data

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/PerforMance308/test2/util"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockEntry(t *testing.T) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)

	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/block.json", GetBlockEntrys())
	assert.NoError(err)

	assert.Equal(GetBlockEntry(1), Block{1, 10})
	assert.Equal(GetBlockEntry(18), Block{18, 10})
}

func TestGetBlockEntrys(t *testing.T) {
	assert := assert.New(t)

	err := util.ParseJSONFile("../configs/block.json", GetBlockEntrys())
	assert.NoError(err)

	assert.Equal(len(GetBlockEntrys().Entrys), 18)
}
