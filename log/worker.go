package log

import (
	"github.com/project-ginza/notifications/const/log_level"
	log "github.com/sirupsen/logrus"
)

func (l *Log) LogWorker() {
	for msg := range l.LogEvents() {
		switch msg.Level {
		case log_level.DEBUG:
			log.Debug(msg.Text)
		case log_level.INFO:
			log.Info(msg.Text)
		case log_level.WARN:
			log.Warn(msg.Text)
		case log_level.FATAL:
			log.Fatal(msg.Text)
		case log_level.ERROR:
			log.Error(msg.Text)
		case log_level.PANIC:
			log.Panic(msg.Text)
		case log_level.TRACE:
			log.Trace(msg.Text)
		}
	}
}
