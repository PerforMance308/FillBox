package fillboxes

import (
	"errors"
	"math/rand"
	"time"

	"github.com/PerforMance308/test2/data"
	"github.com/sirupsen/logrus"
)

type boxes struct {
	filledInterval   []int
	unfilledInterval []*interval
}

// StartFillBlocks start filling the box
func StartFillBlocks() {
	box := &boxes{}

	// init an empty box
	box.init()

	logrus.Infoln("Start!")
	box.start()
}

func (b *boxes) init() {
	intervalCfgs := data.GetIntervalEntrys().Entrys

	b.unfilledInterval = make([]*interval, 0, len(intervalCfgs))

	// get all intervals and create empty intervals asign to box
	for _, i := range intervalCfgs {
		interval := &interval{id: i.ID}
		interval.init()
		b.unfilledInterval = append(b.unfilledInterval, interval)
	}
}

func (b *boxes) start() {
	// is is done, end and output
	if b.isDone() {
		logrus.Infoln("Done!")
		return
	}

	// first step is pick up an interval
	interval, err := b.pickInterval()
	if err != nil {
		logrus.Errorf("error pick interval: %s", err)
		return
	}

	logrus.Infof("pick up interval id: %d", interval.id)

	// second step fill this interval until it is filled
	b.startFillInterval(interval, false)
}

func (b *boxes) startFillInterval(interval *interval, lastPos bool) {
	if interval.isDone() {
		logrus.Infof("Interval Id: %d has been Done!", interval.id)
		b.start()
		return
	}

	// pick next block base on last block position
	blockID, pos := b.picknext(lastPos, interval)
	logrus.Infof("pick up block id: %d", blockID)

	// fill all the intervals that contains this block id
	b.fillAllIntervals(blockID)
	// start next iteration
	b.startFillInterval(interval, pos)
}

func (b *boxes) picknext(lastPos bool, interval *interval) (int, bool) {
	if lastPos {
		// if last position was in the interval, this time random within the whole box
		inorout := inorout()
		if inorout {
			// pick up within the interval
			blockID, err := interval.pickBlock()
			if err != nil {
				logrus.Errorf("error pick block within interval: %s", err)
			}
			return blockID, true
		}

		// pick up within the whole box
		blockID, err := b.pickBlock()
		if err != nil {
			logrus.Errorf("error pick block out of interval: %s", err)
		}
		return blockID, false
	}

	// if last position was not in the interval, this block must be selected from the interval
	blockID, err := interval.pickBlock()
	if err != nil {
		logrus.Errorf("error pick block within interval: %s", err)
	}
	return blockID, true
}

func (b *boxes) fillAllIntervals(blockID int) {
	// iterate all the unfilled intervals, and fill the one contains this blockID
	for index, i := range b.unfilledInterval {
		i.fillBlock(blockID)
		if i.isDone() {
			// if this interval has been done, update box
			b.filledInterval = append(b.filledInterval, i.id)
			b.unfilledInterval = append(b.unfilledInterval[:index], b.unfilledInterval[index+1:]...)
		}
	}
}

func (b *boxes) isDone() bool {
	return len(b.unfilledInterval) == 0
}

func (b *boxes) pickInterval() (*interval, error) {
	// can't pick the last one if still has unfilled interval
	if len(b.unfilledInterval) <= 1 {
		return b.unfilledInterval[0], nil
	}

	var intervals []*interval

	// get all the unfilled intervals except the last one
	for _, i := range b.unfilledInterval {
		intervalCfg := data.GetIntervalEntry(i.id)
		if !intervalCfg.IsLast {
			intervals = append(intervals, i)
		}
	}
	return randomInterval(intervals)
}

func (b *boxes) pickBlock() (int, error) {
	var unfilledBlocks []int
	// get all unfilled blocks
	for _, i := range b.unfilledInterval {
		unfilledBlocks = append(unfilledBlocks, i.unfilledBlocks...)
	}

	return pickBlock(unfilledBlocks)
}

// random where to pick the next block, rate should between 100
func inorout() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand := r.Intn(100)

	rate := data.GetCommonEntry("rate").Value
	return rand < rate
}

// random an unfilled interval, each interval has individual rate, actual rate should be rate/sum([rates...])
func randomInterval(intervals []*interval) (*interval, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	totalRate := 0
	for _, i := range intervals {
		intervalCfg := data.GetIntervalEntry(i.id)
		totalRate += intervalCfg.Rate
	}

	rand := r.Intn(totalRate)
	start := 0
	var end int
	for _, i := range intervals {
		intervalCfg := data.GetIntervalEntry(i.id)
		end += intervalCfg.Rate
		if start <= rand && end > rand {
			return i, nil
		}
		start = end
	}
	return nil, errors.New("error rate")
}
