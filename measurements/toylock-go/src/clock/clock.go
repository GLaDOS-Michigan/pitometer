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
	instant      time.Time // Time of this TimePoint instance
	duration     time.Duration
}

// timePointNow creates a new TimePoint object with the current time
func timePointNow(id int, event EventType, name string) *TimePoint {
	return &TimePoint{
		id:           id,
		event:        event,
		functionName: name,
		instant:      time.Now(),
		duration:     time.Duration(0)}
}

/*****************************************************************************************
************************************** EventLog ******************************************
*****************************************************************************************/

// EventLog represents a sequence of TimePoints
type EventLog struct {
	log    *[]*TimePoint
	nextID int
}

// NewEventLog generates a new log with the specified initial capacity
func NewEventLog(n uint) *EventLog {
	var res = make([]*TimePoint, 0, n)
	return &EventLog{&res, 0}
}

// LogStartEvent adds a new start event to the log
func (el *EventLog) LogStartEvent(name string) {
	var tp = timePointNow(el.nextID, Start, name)
	var newlog = append(*el.log, tp)
	el.log = &newlog
}

// LogEndEvent adds a new end event to the log
func (el *EventLog) LogEndEvent(name string) {
	var tp = timePointNow(el.nextID, End, name)
	el.nextID++
	var previousEvent = (*el.log)[len((*el.log))-1]
	if previousEvent.event != Start {
		fmt.Printf("Error: End event without corresponding start event\n")
		os.Exit(1)
	}
	tp.duration = tp.instant.Sub(previousEvent.instant)
	var newlog = append(*el.log, tp)
	el.log = &newlog
}

// String formats the log into a string
func (el *EventLog) String() string {
	var res = ""
	for _, e := range *el.log {
		var eStr string
		switch e.event {
		case Start:
			eStr = fmt.Sprintf("%d,Start,%s,%v,%v\n", e.id, e.functionName, e.instant, e.duration.Nanoseconds())
		case End:
			eStr = fmt.Sprintf("%d,End,%s,%v,%v\n", e.id, e.functionName, e.instant, e.duration.Nanoseconds())
		default:
			fmt.Printf("Error: Invalid case %v\n", e.event)
		}
		res += eStr
	}
	return res
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
