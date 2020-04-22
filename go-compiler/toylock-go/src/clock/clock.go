package clock

import (
	"fmt"
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
	event        EventType
	functionName string
	instant      time.Time // Time of this TimePoint instance
}

// timePointNow creates a new TimePoint object with the current time
func timePointNow(event EventType, name string) *TimePoint {
	return &TimePoint{
		event:        event,
		functionName: name,
		instant:      time.Now()}
}

/*****************************************************************************************
************************************** EventLog ******************************************
*****************************************************************************************/

// EventLog represents a sequence of TimePoints
type EventLog struct {
	log *[]*TimePoint
}

// NewEventLog generates a new log with the specified initial capacity
func NewEventLog(n uint) *EventLog {
	var res = make([]*TimePoint, 0, n)
	return &EventLog{&res}
}

// LogStartEvent adds a new start event to the log
func (el *EventLog) LogStartEvent(name string) {
	var tp = timePointNow(Start, name)
	var newlog = append(*el.log, tp)
	el.log = &newlog
}

// LogEndEvent adds a new end event to the log
func (el *EventLog) LogEndEvent(name string) {
	var tp = timePointNow(End, name)
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
			eStr = fmt.Sprintf("Start, %s, %v\n", e.functionName, e.instant)
		case End:
			eStr = fmt.Sprintf("End, %s, %v\n", e.functionName, e.instant)
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
	c.count = c.count + 1
}

// GetCount returns the count of the counter
func (c *Counter) GetCount() int {
	return c.count
}
