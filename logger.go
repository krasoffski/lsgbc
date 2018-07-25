package main

import (
	"io"

	"github.com/Sirupsen/logrus"
)

var (
	logger = logrus.Logger{}
)

func initLog(w io.Writer) {
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	logger.Formatter = Formatter
	logger.SetOutput(w)
	logger.SetLevel(logrus.DebugLevel)
}
