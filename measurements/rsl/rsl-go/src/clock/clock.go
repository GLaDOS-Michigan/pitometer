package clock

import (
	"fmt"
	"os"
	"time"
)

/*****************************************************************************************
************************************* TimePoint ******************************************
*****************************************************************************************/

// EventType is Start or End
type EventType int

// Start refers to a start timer event, End refers to an end timer event
const (
	Start EventType = iota
	End
)

// TimePoint contains the information of a clock event
type TimePoint struct {
	id           int
	event        EventType
	functionName string
	instant      time.Duration // Time of this TimePoint instance
}

// timePointNow creates a new TimePoint object
func timePointNow(id int, event EventType, name string, instant time.Duration) *TimePoint {
	return &TimePoint{
		id:           id,
		event:        event,
		functionName: name,
		instant:      instant}
}

/*****************************************************************************************
************************************** Stopwatch *****************************************
*****************************************************************************************/

// Stopwatch represents a sequence of TimePoints
type Stopwatch struct {
	name      string
	initTime  time.Time
	currStart *TimePoint
	nextID    int
}

// NewStopwatch generates a new clean log
func NewStopwatch(name string) *Stopwatch {
	var s = &Stopwatch{
		name:      name,
		initTime:  time.Now(),
		currStart: nil,
		nextID:    0}
	fmt.Printf("stopwatch init,%s,%v\n", s.name, s.initTime)
	return s
}

// LogStartEvent adds a new start event to the log
func (el *Stopwatch) LogStartEvent(name string) {
	var tp = timePointNow(el.nextID, Start, name, time.Since(el.initTime))
	el.currStart = tp
}

// LogEndEvent adds a new end event to the log
func (el *Stopwatch) LogEndEvent(name string) {
	if el.currStart == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	var currStart = el.currStart
	var currEnd = timePointNow(el.nextID, End, name, time.Since(el.initTime))
	fmt.Printf("%d,%s,%v,%v\n", el.currStart.id, el.currStart.functionName, currStart.instant.Nanoseconds(), currEnd.instant.Nanoseconds())
	el.nextID++
	el.currStart = nil
}

// PopStartEvent deletes the last event from the log, which must be a start event
func (el *Stopwatch) PopStartEvent() {
	if el.currStart == nil {
		fmt.Printf("Error: No start time recorded\n")
		os.Exit(1)
	}
	el.currStart = nil
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
