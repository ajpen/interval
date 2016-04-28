// Package interval runs functions at specific intervals, wrapping the standard library ticker 
package interval

import (

	"fmt"
	"time"
	"testing"
)



func TestNewInterval(*testing.T) {
	
	NewInterval(time.Second * 4, func () {fmt.Println("Ran!")})
	time.Sleep(time.Minute * 1)
}


func TestNewTicker(*testing.T) {
	
	var t *time.Ticker

	Ticker(time.Second * 4, func () {fmt.Println("Ran!")}, &t)
	time.Sleep(time.Minute * 1)
	t.Stop()
}