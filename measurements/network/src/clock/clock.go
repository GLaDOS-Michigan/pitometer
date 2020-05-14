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
	startTime time.Time
	log       *[]*TimePoint
	nextID    int
}

// NewStopwatch generates a new log with the specified initial capacity
func NewStopwatch(n uint, name string) *Stopwatch {
	var res = make([]*TimePoint, 0, n)
	return &Stopwatch{
		name:      name,
		startTime: time.Now(),
		log:       &res,
		nextID:    0}
}

// LogStartEvent adds a new start event to the log
func (el *Stopwatch) LogStartEvent(name string) {
	var tp = timePointNow(el.nextID, Start, name, time.Since(el.startTime))
	var newlog = append(*el.log, tp)
	el.log = &newlog
}

// LogEndEvent adds a new end event to the log
func (el *Stopwatch) LogEndEvent(name string) {
	var tp = timePointNow(el.nextID, End, name, time.Since(el.startTime))
	el.nextID++
	var newlog = append(*el.log, tp)
	el.log = &newlog
}

// String formats the log into a string
func (el *Stopwatch) String() string {
	var res = fmt.Sprintf("%s,%v\n", el.name, el.startTime)
	for _, e := range *el.log {
		var eStr string
		switch e.event {
		case Start:
			eStr = fmt.Sprintf("%d,Start,%s,%v\n", e.id, e.functionName, e.instant.Nanoseconds())
		case End:
			eStr = fmt.Sprintf("%d,End,%s,%v\n", e.id, e.functionName, e.instant.Nanoseconds())
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
