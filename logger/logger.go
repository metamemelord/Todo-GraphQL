package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	log.Println("Initialized a custom logger")
	return log.New(os.Stdout, "[Custom logger]: ", 0)
}
