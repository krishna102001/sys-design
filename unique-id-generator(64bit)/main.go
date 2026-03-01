package main

import (
	"log"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////||
//  unsigned	      //  timestamp in  //  DataCenter ID  //   machine ID   //  Sequence ID||
//  interger (1 bit) //	 unix (41 bit) //   (5 bit)       //    (5 bit)     //    (12 bit)  ||
////////////////////////////////////////////////////////////////////////////////////////////||

const (
	machineIDbits    = uint64(5)
	datacenterIDbits = uint64(5)
	sequenceIDbits   = uint64(12)

	max_datacenterID uint = uint64(-1) << (3294)
	max_machineID    uint64
	max_sequenceID   uint64
	tw_epoch         = int64(1288834974657)
)

type worker struct {
	mu           sync.RWMutex
	datacenterID uint64
	machineID    uint64
	sequenceID   uint64
}

func New(dID, mID, sID uint64) *worker {
	return &worker{
		datacenterID: dID,
		machineID:    mID,
		sequenceID:   sID,
	}
}

func (w *worker) get_time_stamp_now() uint64 {
	return uint64(time.Now().UnixMilli())
}

func main() {
	timeepoch := time.UnixMilli(int64(tw_epoch)).UTC()
	log.Println("timeepoch", timeepoch)
}

// XOR (^) operator - it return 1 if the bits are different and 0 if they are same
// leftshift (<<) operator - it shift the bit of binary number toward left by specified positions and replace empty spaces with 0 from the right. 0010 << 1 = 0100
// bitwise OR (|) operator - it return 1 if anyone bit is 1 else 0 if both bit are zero
// bitwise AND (&) operator - it return 1 if both bit are 1 else it return 0
