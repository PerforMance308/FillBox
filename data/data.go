package data

import (
	"github.com/PerforMance308/test2/util"
	"github.com/sirupsen/logrus"
)

// InitConfig all config data
func InitConfig() {
	err := util.ParseJSONFile("configs/block.json", GetBlockEntrys())

	if err != nil {
		logrus.Panicf("Failed to load json file: block.json, err: %s", err)
	}

	err = util.ParseJSONFile("configs/interval.json", GetIntervalEntrys())
	if err != nil {
		logrus.Panicf("Failed to load json file: interval.json, err: %s", err)
	}

	err = util.ParseJSONFile("configs/common.json", GetCommonEntrys())
	if err != nil {
		logrus.Panicf("Failed to load json file: common.json, err: %s", err)
	}
}
