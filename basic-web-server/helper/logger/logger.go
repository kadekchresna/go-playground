package logger

import "go.uber.org/zap"

var (
	cfg zap.Config
)

type LoggerConfig struct {
	Level string
}

type Logger struct {
	Log *zap.Logger
}

func Log(cfg LoggerConfig) Logger {

	config := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		ErrorOutputPaths: []string{"stderr"},
	}
	logger := zap.Must(config.Build())
	defer logger.Sync()

	return Logger{Log: logger}

}
