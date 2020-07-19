package data

// Block config declare
type Block struct {
	ID   int `json:"id"`
	Rate int `json:"rate"`
}

// BlockEntryMap stores Block entrys
type BlockEntryMap map[int]Block

// Blocks stores Blocks and entry map
type Blocks struct {
	Entrys []Block `json:"block"`
	entrys BlockEntryMap
}

var blockEntrys = Blocks{entrys: make(BlockEntryMap)}

// InitMap inits block entry map
func (bs *Blocks) InitMap() {
	for _, entry := range bs.Entrys {
		bs.entrys[entry.ID] = entry
	}
}

// GetBlockEntrys returns Blocks point
func GetBlockEntrys() *Blocks {
	return &blockEntrys
}

// GetBlockEntry returns all blocks entriy
func GetBlockEntry(id int) Block {
	return blockEntrys.entrys[id]
}
