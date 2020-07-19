package main

import (
	"github.com/PerforMance308/test2/data"
	"github.com/PerforMance308/test2/fillboxes"
	"github.com/PerforMance308/test2/options"
)

func main() {
	// init options
	options.Init("logging")
	// init configs
	data.InitConfig()

	// main logic
	fillboxes.StartFillBlocks()
}
