package clock

import (
	"fmt"
	"os"
	"time"
)

/*****************************************************************************************
************************************* TimePoint ******************************************
*****************************************************************************************/

// TimeInterval contains the information of a clock event
type TimeInterval struct {
	id           int
	functionName string
	timestamp    time.Time
	startTime    time.Duration //start time of this interval
	endTime      time.Duration //end time of this interval
}

// newTimeInterval creates a new TimeInterval object with void start and end times
func newTimeInterval(id int, name string) *TimeInterval {
	return &TimeInterval{
		id:           id,
		functionName: name}
}

func (ti *TimeInterval) logStartTime(st time.Duration) {
	ti.startTime = st
}

func (ti *TimeInterval) logEndTime(et time.Duration) {
	ti.endTime = et
	ti.timestamp = time.Now()
}

// PrintTimeInterval prints the time interval
func (ti *TimeInterval) PrintTimeInterval() {
	fmt.Printf("%d,%s,%v,%v,%s\n", ti.id, ti.functionName, ti.startTime.Nanoseconds(), ti.endTime.Nanoseconds(), ti.timestamp.String())
}

/*****************************************************************************************
************************************** Stopwatch *****************************************
*****************************************************************************************/

// Stopwatch represents a sequence of TimePoints
type Stopwatch struct {
	name         string
	initTime     time.Time
	currInterval *TimeInterval
	nextIndex    int
}

// NewStopwatch generates a new log pre-initialized to length n
func NewStopwatch(name string) *Stopwatch {
	var s = &Stopwatch{
		name:         name,
		initTime:     time.Now(),
		currInterval: nil,
		nextIndex:    0}
	fmt.Printf("stopwatch init,%s,%v\n", s.name, s.initTime)
	return s
}

// LogStartEvent adds a temp start event to currInterval in the stopwatch
func (el *Stopwatch) LogStartEvent() {
	if el.currInterval != nil {
		fmt.Printf("Error: Pending interval already present\n")
		os.Exit(1)
	}
	var ti = newTimeInterval(el.nextIndex, "generic event")
	ti.logStartTime(time.Since(el.initTime))
	el.currInterval = ti
}

// LogEndEvent adds a new end event to the log
func (el *Stopwatch) LogEndEvent(name string) {
	if el.currInterval == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	var ti = el.currInterval
	ti.logEndTime(time.Since(el.initTime))
	ti.functionName = name
	ti.PrintTimeInterval()
	el.nextIndex++
	el.currInterval = nil
}

// PopStartEvent deletes the last event from the log, which must be a start event
func (el *Stopwatch) PopStartEvent() {
	if el.currInterval == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	el.currInterval = nil
}

/*****************************************************************************************
************************************** Counter *******************************************
*****************************************************************************************/

// Counter is used to count events in code
type Counter struct {
	count int
	name  string
}

// NewCounter is the constructor for a Counter
func NewCounter(name string) *Counter {
	return &Counter{0, name}
}

// Increment increments the counter by 1
func (c *Counter) Increment() {
	c.count++
}

// GetCount returns the count of the counter
func (c *Counter) GetCount() int {
	return c.count
}
