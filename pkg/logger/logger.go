package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	// ламберджэк нужен для роллящихся файлов логов
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

type LoggerConfig struct {
	Logfile  string
	Loglevel zapcore.Level
}

// для соответствия уровня логгирования в конфиге и внутренним уровнем логера
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[viper.GetString("logger.level")]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func InitLogger() {
	if sugarLogger != nil {
		return
	}

	logFile := viper.GetString("logger.logfile")
	logLevel := getLoggerLevel()

	var logWriter zapcore.WriteSyncer

	if logFile != "stdout" {
		logWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename: logFile,
			//TODO: подумать как достать из конфига
			MaxSize:   1 << 30, //1G
			LocalTime: true,
			Compress:  true,
		})
	} else {
		logWriter = zapcore.AddSync(os.Stdout)
	}

	//encoder := zap.NewProductionEncoderConfig()
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

func InitLoggerByConfig(config LoggerConfig) {
	if sugarLogger != nil {
		return
	}

	logFile := config.Logfile
	logLevel := config.Loglevel

	var logWriter zapcore.WriteSyncer

	if logFile != "stdout" {
		logWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename: logFile,
			//TODO: подумать как достать из конфига
			MaxSize:   1 << 30, //1G
			LocalTime: true,
			Compress:  true,
		})
	} else {
		logWriter = zapcore.AddSync(os.Stdout)
	}

	//encoder := zap.NewProductionEncoderConfig()
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

// набор методов нашего логгера для использования извне

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
