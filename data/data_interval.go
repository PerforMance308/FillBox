package data

// Interval config declare
type Interval struct {
	ID       int   `json:"id"`
	BlockIds []int `json:"blocks"`
	IsLast   bool  `json:"islast"`
	Rate     int   `json:"rate"`
}

// IntervalEntryMap stores Interval entrys
type IntervalEntryMap map[int]Interval

// Intervals stores Intervals and entry map
type Intervals struct {
	Entrys []Interval `json:"interval"`
	entrys IntervalEntryMap
	last   Interval
}

var intervalEntrys = Intervals{entrys: make(IntervalEntryMap)}

// InitMap inits Interval entry map
func (is *Intervals) InitMap() {
	for _, entry := range is.Entrys {
		is.entrys[entry.ID] = entry
		if entry.IsLast {
			is.last = entry
		}
	}
}

// GetIntervalEntrys returns Intervals point
func GetIntervalEntrys() *Intervals {
	return &intervalEntrys
}

// GetIntervalEntry returns all Intervals entriy
func GetIntervalEntry(id int) Interval {
	return intervalEntrys.entrys[id]
}

// GetLastInterval returns last Interval
func GetLastInterval() Interval {
	return intervalEntrys.last
}
