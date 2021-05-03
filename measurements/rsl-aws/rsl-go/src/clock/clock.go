package clock

import (
	"fmt"
	"os"
	"time"
)

var ON bool
var INIT_SIZE int

/*****************************************************************************************
************************************* TimePoint ******************************************
*****************************************************************************************/

// TimeInterval contains the information of a clock event
type TimeInterval struct {
	id           int
	functionName string
	startTime    time.Duration //start time of this interval
	endTime      time.Duration //end time of this interval
}

// newTimeInterval creates a new TimeInterval object with void start and end times
func newTimeInterval(id int, name string) *TimeInterval {
	return &TimeInterval{
		id:           id,
		functionName: name,
		startTime:    0,
		endTime:      0}
}

func (ti *TimeInterval) logStartTime(st time.Duration) {
	ti.startTime = st
}

func (ti *TimeInterval) logEndTime(et time.Duration) {
	ti.endTime = et
}

/*****************************************************************************************
************************************** Stopwatch *****************************************
*****************************************************************************************/

// Stopwatch represents a sequence of TimePoints
type Stopwatch struct {
	name         string
	initTime     time.Time
	log          *[]TimeInterval
	currInterval *TimeInterval
	nextIndex    int
}

// NewStopwatch generates a new log pre-initialized to length n
func NewStopwatch(name string) *Stopwatch {
	var l = make([]TimeInterval, INIT_SIZE, INIT_SIZE)
	var s = &Stopwatch{
		name:         name,
		initTime:     time.Now(),
		log:          &l,
		currInterval: nil,
		nextIndex:    0}
	fmt.Printf("stopwatch init,%s,%v\n", s.name, s.initTime)
	return s
}

// LogStartEvent adds a temp start event to currInterval in the stopwatch
func (el *Stopwatch) LogStartEvent(name string) {
	if !ON {
		return
	}
	if el.currInterval != nil {
		fmt.Printf("Error: Pending interval already present\n")
		os.Exit(1)
	}
	var st = time.Since(el.initTime)
	var ti = newTimeInterval(el.nextIndex, name)
	ti.logStartTime(st)
	el.currInterval = ti
}

// LogEndEvent adds a new end event to the log
func (el *Stopwatch) LogEndEvent(name string) {
	if !ON {
		return
	}
	if el.currInterval == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	var ti = el.currInterval
	if el.nextIndex >= cap(*el.log) {
		ti.logEndTime(time.Since(el.initTime))
		el.nextIndex++
		el.currInterval = nil
		return
	}
	el.currInterval = nil
	var et = time.Since(el.initTime)
	ti.logEndTime(et)
	(*el.log)[el.nextIndex] = *ti
	el.nextIndex++
}

// PopStartEvent deletes the last event from the log, which must be a start event
func (el *Stopwatch) PopStartEvent() {
	if !ON {
		return
	}
	if el.currInterval == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	el.currInterval = nil
}

func (el *Stopwatch) MakeStartEvent() *TimeInterval {
	if !ON {
		return nil
	}
	var st = time.Since(el.initTime)
	var ti = newTimeInterval(0, el.name)
	ti.logStartTime(st)
	return ti
}

func (el *Stopwatch) RecordEndEvent(ti *TimeInterval) {
	if !ON {
		return
	}
	if el.nextIndex >= cap(*el.log) {
		ti.logEndTime(time.Since(el.initTime))
		el.nextIndex++
		el.currInterval = nil
		return
	}
	ti.id = el.nextIndex
	var et = time.Since(el.initTime)
	ti.logEndTime(et)
	(*el.log)[el.nextIndex] = *ti
	el.nextIndex++
	el.currInterval = nil
}

// PrintLog prints the log line by line
func (el *Stopwatch) PrintLog() {
	fmt.Printf("%s,%v\n", el.name, el.initTime)
	for i, ti := range *el.log {
		if i >= el.nextIndex {
			break
		}
		fmt.Printf("%d,%s,%v,%v\n", ti.id, ti.functionName, ti.startTime.Nanoseconds(), ti.endTime.Nanoseconds())
	}
	fmt.Printf("End of log %s\n", el.name)
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
