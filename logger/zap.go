package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"kube-x/global"
)

func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

// getEncoder returns a zapcore.Encoder
// log encoder is used to encode the log message
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter returns a zapcore.WriteSyncer
// log writer is used to write the log message to the file
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.CONF.Logger.Filename,   // log file path
		MaxSize:    global.CONF.Logger.MaxSize,    // megabytes
		MaxBackups: global.CONF.Logger.MaxBackups, // number of backups
		MaxAge:     global.CONF.Logger.MaxAge,     // days
		Compress:   global.CONF.Logger.Compress,   // disabled by default
	}
	return zapcore.AddSync(lumberJackLogger)
}
