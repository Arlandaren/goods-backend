package main

import (
	"log"
	"os"
)

var Logger *log.Logger

type LoggerWrapper struct {
	*log.Logger
}

func NewLogger() *LoggerWrapper {
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		Logger.Println("Не удалось создать файл логов: ", err)
	}
	logger := log.New(file, "|||LOG|||", log.Ldate|log.Ltime|log.Lshortfile)
	return &LoggerWrapper{
		Logger: logger,
	}
}

func (lw *LoggerWrapper) Info(message string) {
	lw.Println("[INFO]", message)
}

func (lw *LoggerWrapper) Error(message string) {
	lw.Println("[ERROR]", message)
}
