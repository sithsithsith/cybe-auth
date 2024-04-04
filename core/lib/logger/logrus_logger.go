package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type GinFormatter struct {
	// default timestamp format = "Jan _2 15:04:05.000"
	TimestampFormat string
	// hide keys in field
	HideKeys bool
	//no space between fields
	NoFieldsSpace bool
	//no enforcing colors
	NoColors bool
}

func (f *GinFormatter) writeField(b *bytes.Buffer, entry *logrus.Entry, field string) {
	if f.HideKeys {
		fmt.Fprintf(b, "[%v]", entry.Data[field])
	} else {
		fmt.Fprintf(b, "[%s:%v]", field, entry.Data[field])
	}
	if !f.NoFieldsSpace {
		b.WriteString(" ")
	}
}

func (f *GinFormatter) getColorLevel(entry *logrus.Entry) int {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 31 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return levelColor
}

func (f *GinFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := f.getColorLevel(entry)
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func NewLogger(lvl logrus.Level) *logrus.Logger {
	f, _ := os.OpenFile("logs/http-logs.txt", os.O_CREATE|os.O_WRONLY, 0777)
	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stderr, f))
	logger.SetLevel(lvl)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	},
	)
	return logger
}
