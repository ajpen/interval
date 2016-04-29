# Interval

Executes a function periodically. 


## Description

Interval is a wrapper for the ticker and tick functions in the time library. It wraps everything into a single function call.
The function passed is executed in go routine and should prepared as such. All intervals started will spawn a goroutine; 
if you wish to execute multiple functions at the same intervals, you should wrap the calls into a single function and pass it.
Make sure the calls do not block or else the preceeding calls will be delayed.


## Install

`go get github.com/ajpen/interval`


## Example

`NewInterval(time.Second * 4, func () {fmt.Println("Ran!")})`
