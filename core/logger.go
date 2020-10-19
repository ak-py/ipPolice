package core

import (
	"fmt"
	log_rotation "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

var ipPoliceLogger *log.Logger

func Logger() *log.Logger {
	if ipPoliceLogger == nil {
		ipPoliceLogger = newLogger()
	}
	return ipPoliceLogger
}

func newLogger() *log.Logger {
	logsDirPath := logsDirectory()

	baseFileName := "ip-police" + ".log"
	fileName := filepath.FromSlash(logsDirPath + PathSeparator + baseFileName)
	fmt.Println(fileName)

	mw := io.MultiWriter(os.Stdout, &log_rotation.Logger{
		Filename:   fileName,
		MaxSize:    2,
		MaxAge:     365,
		MaxBackups: 10,
	})

	baseLogger := log.New(mw, "ip-police: ", log.LstdFlags|log.Lshortfile)
	return baseLogger
}

// Regex to extract just the function name (and not the module path)
var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

func callerFuncName(skip int) string {
	fnName := "<unknown>"
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	}
	return fnName
}
