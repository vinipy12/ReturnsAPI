package logger

import "log"

type Logger struct {
	messageQueue chan string
}

func NewLogger() *Logger {
	return &Logger{
		messageQueue: make(chan string, 1000),
	}
}

func (l *Logger) InitLogger() {
	go func() {
		for msg := range l.messageQueue {
			logOutput := msg
			log.Print(logOutput)
		}
	}()
}

func (l *Logger) Log(message string) {
	l.messageQueue <- message
}
