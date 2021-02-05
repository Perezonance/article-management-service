package log

import (
	"fmt"
	"time"
)

const (
	timeFormat = "01/02/2006 15:04:05"

	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorCyan   = "\033[32m" //32 - green
	colorYellow = "\033[33m"
)

//InfoLog provides a info level structured log
func InfoLog(msg string) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf(colorCyan+"%v::INFO LOG  >> "+colorReset+msg+"\n", currTime)
}

//ErrorLog provides an error level structured log
func ErrorLog(msg string, err error) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf(colorRed+"%v::ERROR LOG >> "+colorReset+msg+":\n%v\n", currTime, err)
}

//DebugLog provides a debug level structured log
func DebugLog(msg string) {
	currTime := time.Now().Format(timeFormat)
	fmt.Printf(colorYellow+"%v::DEBUG LOG >> "+colorReset+msg+"\n", currTime)
}
