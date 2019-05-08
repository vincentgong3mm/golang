package doriginal

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	filename string
	*log.Logger
}

var logger *Logger
var once sync.Once

func GetLogInstance() *Logger {
	once.Do(func() {
		logger = createLogger("dom.log")
	})
	return logger
}

func createLogger(fname string) *Logger {
	file, _ := os.OpenFile("dom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	return &Logger{filename: fname,
		Logger: log.New(file, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)}
}

func init() {
	fmt.Println("import log")
}
