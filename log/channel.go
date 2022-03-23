package log

import "github.com/project-ginza/notifications/const/log_level"

var LoggingCh *Log

func createLogChannel() *Log {
	return &Log{
		Channel: make(chan Message, 10000),
	}
}

func (l *Log) LogEvents() <-chan Message {
	return l.Channel
}

func (l *Log) PushDebugMessageToChannel(msg string) {
	l.Channel <- Message{log_level.DEBUG, msg}
}

func (l *Log) PushInfoMessageToChannel(msg string) {
	l.Channel <- Message{log_level.INFO, msg}
}

func (l *Log) PushWarnMessageToChannel(msg string) {
	l.Channel <- Message{log_level.WARN, msg}
}

func (l *Log) PushFatalMessageToChannel(msg string) {
	l.Channel <- Message{log_level.FATAL, msg}
}

func (l *Log) PushErrorMessageToChannel(msg string) {
	l.Channel <- Message{log_level.ERROR, msg}
}

func (l *Log) PushPanicMessageToChannel(msg string) {
	l.Channel <- Message{log_level.PANIC, msg}
}

func (l *Log) PushTraceMessageToChannel(msg string) {
	l.Channel <- Message{log_level.TRACE, msg}
}

func init() {
	LoggingCh = createLogChannel()
}
