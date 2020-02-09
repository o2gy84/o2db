package config

type Logger struct {
	Level           *string
	FullTimestamps  *bool
	TimestampFormat *string
	Syslog          *bool
}
