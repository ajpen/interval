// Package interval runs functions at specific intervals, wrapping the standard library ticker 
package interval

import (

	"time"
	"errors"
)



// NewInterval runs the task after the duration, forever. 
func NewInterval(duration time.Duration, task func()) error {
	if duration <= 0 {
		return errors.New("Duration must be greater than 0")
	}

	ticks := time.Tick(duration)

	go func (task func()) {
			
		for _ = range ticks {
			task()
		}
	}(task)

	return nil
}