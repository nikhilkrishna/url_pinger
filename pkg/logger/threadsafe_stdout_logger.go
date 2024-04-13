package logger

import (
	"fmt"
	"sync"
)


type StdoutLogger struct {
	logChan chan string 
	wg *sync.WaitGroup
}


func NewStdoutLogger() Logger {  
    logger := &StdoutLogger{
        logChan: make(chan string, 100),
        wg:      new(sync.WaitGroup),
    }
    logger.wg.Add(1)
    go logger.processLogs()
    return logger
}

func (l *StdoutLogger) Log(message string) {
    l.logChan <- message
}


func (l *StdoutLogger) Close() {
    close(l.logChan)
    l.wg.Wait()  
}

func (l *StdoutLogger) processLogs() {
    defer l.wg.Done()
    for msg := range l.logChan {
        fmt.Println(msg)
    }
}