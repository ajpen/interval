// Package interval runs functions at specific intervals, wrapping the standard library ticker 
package interval

import (

	"time"
	"errors"
)



// NewInterval runs the task after the duration, forever. Tasks should be designed as goroutine compatible
func NewInterval(duration time.Duration, task func()) error {
	if duration <= 0 {
		return errors.New("Duration must be greater than 0")
	}

	ticks := time.Tick(duration)

	go func (task func()) {
			
		for _ = range ticks {
			go task()
		}
	}(task)

	return nil
}


// NewTicker runs the task after the duration. The underlying ticker is exposed through ticker
func Ticker(duration time.Duration, task func(), ticker **time.Ticker) error {
	
	if duration <= 0 {
		return errors.New("Duration must be greater than 0")
	}

	*ticker = time.NewTicker(duration)

	go func (task func()) {
		
		for _ = range (*ticker).C {

			go task()
		}
	}(task)

	return nil
}