package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.Logger

func init() {

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}

	logFile, errFile := os.OpenFile("logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errFile != nil {
		panic(errFile)
	}

	writeSyncer := zapcore.AddSync(logFile)
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config.EncoderConfig),
		writeSyncer,
		zap.InfoLevel,
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		consoleWriteSyncer,
		zap.InfoLevel,
	)
	core := zapcore.NewTee(fileCore, consoleCore)
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Warn(v.Error(), fields...)
	case string:
		log.Warn(v, fields...)
	}
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
