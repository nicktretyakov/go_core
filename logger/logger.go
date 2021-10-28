package logger

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (Logger *log.Logger)

func SetupLogger() {

	y, m, d := time.Now().Date()

	name := "logs-" + strconv.Itoa(y) + "." + strconv.Itoa(int(m)) + "." + strconv.Itoa(d) +".txt"

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	Logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		log.Fatal(err)
	}

	Logger.SetOutput(file)
}
