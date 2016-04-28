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