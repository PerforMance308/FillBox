package fillboxes

import (
	"errors"
	"math/rand"
	"time"

	"github.com/PerforMance308/test2/data"
)

type interval struct {
	id             int
	filledBlocks   []int
	unfilledBlocks []int
}

// init an empty interval
func (i *interval) init() {
	intervalCfgs := data.GetIntervalEntry(i.id)
	i.unfilledBlocks = intervalCfgs.BlockIds
}

// pick a block within this interval
func (i *interval) pickBlock() (int, error) {
	return pickBlock(i.unfilledBlocks)
}

// fill block by blockid in this interval
func (i *interval) fillBlock(blockID int) {
	for index, b := range i.unfilledBlocks {
		if b == blockID {
			i.filledBlocks = append(i.filledBlocks, blockID)
			i.unfilledBlocks = append(i.unfilledBlocks[:index], i.unfilledBlocks[index+1:]...)
			break
		}
	}
}

func (i *interval) isDone() bool {
	return len(i.unfilledBlocks) == 0
}

// random an unfilled block, each block has individual rate, actual rate should be rate/sum([rates...])
func pickBlock(blockIDs []int) (int, error) {
	totalRate := 0
	for _, blockID := range blockIDs {
		blockCfg := data.GetBlockEntry(blockID)
		totalRate += blockCfg.Rate
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand := r.Intn(totalRate)

	start := 0
	var end int

	for _, blockID := range blockIDs {
		blockCfg := data.GetBlockEntry(blockID)
		end += blockCfg.Rate
		if start <= rand && end > rand {
			return blockID, nil
		}
		start = end
	}
	return 0, errors.New("error rate")
}
