package log

type Message struct {
	Level int
	Text  string
}

type Log struct {
	Channel chan Message
}

func FormLog(logType int, msg string) Message {
	return Message{logType, msg}
}
