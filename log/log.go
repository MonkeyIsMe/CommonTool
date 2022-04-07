package log

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

var logFile *os.File

func init() {
	var err error
	logFile, err = os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

}

func ErrorContextf(ctx context.Context, format string) {
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logInfo := fmt.Sprintf("[Error]:%+v:%s", ctx, format)
	log.Println(logInfo)

}

func DebugContextf(ctx context.Context, format string) {
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logInfo := fmt.Sprintf("[Debug]:%+v:%s", ctx, format)
	log.Println(logInfo)

}

func InfoContextf(ctx context.Context, format string) {
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	logInfo := fmt.Sprintf("[Info]:%+v:%s", ctx, format)
	log.Println(logInfo)

}
