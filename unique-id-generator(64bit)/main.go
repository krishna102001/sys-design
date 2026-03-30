package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////||
//  unsigned	      //  timestamp in  //  DataCenter ID  //   machine ID   //  Sequence ID||
//  interger (1 bit) //	 unix (41 bit) //   (5 bit)       //    (5 bit)     //    (12 bit)  ||
////////////////////////////////////////////////////////////////////////////////////////////||

const (
	machineIDbits    = uint64(5)  // Binary: 00101
	datacenterIDbits = uint64(5)  // Binary: 00101
	sequenceIDbits   = uint64(12) // Binary: 01100

	max_datacenterID = (int64(-1)) ^ (int64(-1) << datacenterIDbits)
	max_machineID    = (int64(-1)) ^ (int64(-1) << machineIDbits)
	max_sequenceID   = int64(-1) ^ (int64(-1) << sequenceIDbits)

	timeLeft     = uint8(22)
	dataLeft     = uint8(17)
	sequenceLeft = uint8(12)
	tw_epoch     = int64(1288834974657)
)

type worker struct {
	mu            sync.Mutex
	lastTimeStamp int64 // record the last time stamp when id was generated
	datacenterID  int64
	machineID     int64
	sequenceID    int64
}

func New(dID, mID int64) *worker {
	return &worker{
		lastTimeStamp: 0,
		datacenterID:  dID,
		machineID:     mID,
		sequenceID:    0,
	}
}

func (w *worker) get_time_stamp_now() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *worker) NextUniqueID() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.generateUniqueID()
}

func (w *worker) generateUniqueID() (uint64, error) {
	timeStamp := w.get_time_stamp_now()
	log.Printf("timestamp i got %v", timeStamp)
	if timeStamp < w.lastTimeStamp {
		return 0, errors.New("time is moving backwards,waiting until")
	}
	if w.lastTimeStamp == timeStamp {
		w.sequenceID = (w.sequenceID + 1) & max_sequenceID
		log.Printf("sequence id %v", w.sequenceID)
		if w.sequenceID == 0 {
			for timeStamp <= w.lastTimeStamp {
				timeStamp = w.get_time_stamp_now()
			}
		}
	} else {
		w.sequenceID = 0
	}

	w.lastTimeStamp = timeStamp
	id := ((timeStamp - tw_epoch) << timeLeft) | (w.datacenterID << int64(dataLeft)) | (w.machineID << int64(sequenceLeft)) | w.sequenceID

	return uint64(id), nil
}

func main() {
	var wg sync.WaitGroup
	w := New(5, 5)

	ch := make(chan uint64, 10000)
	count := 10000
	wg.Add(count)
	defer close(ch)
	// Concurrently count goroutines for snowFlake ID generation
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			id, _ := w.NextUniqueID()
			ch <- id
		}()
	}
	wg.Wait()
	m := make(map[uint64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// If there is a key with id in the map, it means that the generated snowflake ID is duplicated
		_, ok := m[id]
		if ok {
			fmt.Printf("repeat id %d\n", id)
			return
		}
		// store id as key in map
		m[id] = i
		// log.Printf("unique id are %d \n", id)

	}
	// successfully generated snowflake ID
	fmt.Println("All", len(m), "snowflake ID Get successed!")

}

// XOR (^) operator - it return 1 if the bits are different and 0 if they are same
// leftshift (<<) operator - it shift the bit of binary number toward left by specified positions and replace empty spaces with 0 from the right. 0010 << 1 = 0100
// bitwise OR (|) operator - it return 1 if anyone bit is 1 else 0 if both bit are zero
// bitwise AND (&) operator - it return 1 if both bit are 1 else it return 0
// %b - it is used to print the binary number
