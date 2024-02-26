package logic

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/satyarth42/id-gen/config"
)

var counter atomic.Int64
var lastTimeStamp atomic.Int64
var baseID int64

func init() {
	conf := config.GetConfig()

	baseID = 0
	baseID = setDCBits(baseID, conf.DC)
	baseID = setServerBits(baseID, conf.Server)

	counter.Store(0)
	lastTimeStamp.Store(0)
}

func GenerateID() (int64, error) {
	var id int64 = baseID
	var err error
	id, err = setMillisecondAndCounterBits(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func setMillisecondAndCounterBits(id int64) (int64, error) {
	currentTime := time.Now()
	epochMs := currentTime.UnixMilli()

	var counterValue int64
	if lastTimeStamp.Load() == epochMs {
		counterValue = counter.Add(1)
	} else {
		counter.Store(0)
		counterValue = 0
		lastTimeStamp.Store(epochMs)
	}

	if counterValue >= (1 << 12) {
		return 0, fmt.Errorf("counter overflow at time: %s", currentTime)
	}

	lastTimeStamp.Store(epochMs)

	timeMask := int64(1<<41) - 1
	msBits := epochMs & timeMask

	counterMask := int64(1<<12) - 1

	counterBits := counterValue & counterMask

	id = id | (msBits << 22) | (counterBits)

	return id, nil
}

func setDCBits(id int64, dc int8) int64 {
	mask := getMaskOfBits(5)

	dcBits := int64(dc) & mask

	id = id | (dcBits << 17)

	return id
}

func setServerBits(id int64, server int8) int64 {
	mask := getMaskOfBits(5)

	serverBits := int64(server) & mask

	id = id | (serverBits << 12)

	return id
}
