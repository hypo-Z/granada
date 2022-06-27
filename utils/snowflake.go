package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	workerIDOffset  = 12
	timestampOffset = 22
	maxTimestamp    = 1<<41 - 1
	maxWorkerID     = 1<<10 - 1
	maxSequence     = 1<<12 - 1
)

var (
	// ErrInvalidWorkerID indicates that worker id is out of valid bounds
	ErrInvalidWorkerID = errors.New("worker ID is not within bounds")
	since              = time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).UnixNano()
)

type Snowflake <-chan int64

// New constructs generator for snowflake IDs
// ErrInvalidWorkerID is returned if WorkerID is not within [0, 1024)
func New(workerID uint64) (Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, ErrInvalidWorkerID
	}

	sf := make(chan int64)
	go generator(workerID, sf)
	return sf, nil
}

func generator(workerID uint64, c chan<- int64) {
	last := timestamp()
	seq := uint64(0)
	for {
		ts := timestamp()
		if ts < last {
			ts = nextMillisec(last)
		}

		if ts != last {
			seq = 0
			last = ts
		} else if seq == maxSequence {
			ts = nextMillisec(ts)
			seq = 0
			last = ts
		} else {
			seq++
		}

		id := int64((ts << timestampOffset) | (workerID << workerIDOffset) | seq)
		c <- id
	}
}

func nextMillisec(ts uint64) uint64 {
	i := timestamp()
	for ; i <= ts; i = timestamp() {
		time.Sleep(100 * time.Microsecond)
	}
	return i
}

func timestamp() uint64 {
	return (uint64(time.Now().UnixNano()-since) / uint64(time.Millisecond)) & maxTimestamp
}

// GetID 生成ID
func GetID() string {
	id, err := New(42)
	if err != nil {
		fmt.Printf("GetID error: %v\n", err)
	}
	uid := strconv.FormatInt(<-id, 10)
	return uid
}
