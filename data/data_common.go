package data

// Common config declare
type Common struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// CommonEntryMap stores Common entrys
type CommonEntryMap map[string]Common

// Commons stores Commons and entry map
type Commons struct {
	Entrys []Common `json:"common"`
	entrys CommonEntryMap
}

var commonEntrys = Commons{entrys: make(CommonEntryMap)}

// InitMap inits Common entry map
func (is *Commons) InitMap() {
	for _, entry := range is.Entrys {
		is.entrys[entry.Key] = entry
	}
}

// GetCommonEntrys returns commonEntrys object
func GetCommonEntrys() *Commons {
	return &commonEntrys
}

// GetCommonEntry returns all Common configs entriy
func GetCommonEntry(key string) Common {
	return commonEntrys.entrys[key]
}
