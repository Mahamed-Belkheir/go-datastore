package logging

type LogEvent struct {
	level string
	message string
}

func New(level, message string) LogEvent {
	return LogEvent{
		level: level,
		message: message,
	}
}