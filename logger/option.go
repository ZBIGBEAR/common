package logger

func WithFeiShuNotify(webhook string) Option {
	return func(l *log) {
		l.notify = newFeiShuNotify(webhook)
	}
}

func WithServiceName(serviceName string) Option {
	return func(l *log) {
		l.serviceName = serviceName
	}
}
