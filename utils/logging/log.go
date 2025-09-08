package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogLevelInfo = "info"
)

func InitLogger(programName, logLevel string) (*zap.Logger, error) {
	return InitLoggerWithLogFile(programName, "", logLevel)
}

func InitLoggerWithLogFile(programName, logFilePath, logLevel string) (log *zap.Logger, err error) {
	if logLevel == "" {
		logLevel = defaultLogLevelInfo
	}

	config := zap.NewProductionConfig()

	// Set log level based on logLevel config
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case defaultLogLevelInfo:
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	config.Level = zap.NewAtomicLevelAt(level)

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.Encoding = "json"

	// Create a file to write logs
	if logFilePath != "" {
		// Ensure the directory exists
		dir := filepath.Dir(logFilePath)
		if err = os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("%s init logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return nil, err
		}

		// Create a file to write logs
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("%s init logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return nil, err
		}
		// Create a core that writes logs to the file
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(file),
			level,
		)
		// Build the logger with the file core
		log = zap.New(fileCore, zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		var err error
		log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			fmt.Printf("%s init logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return nil, err
		}
	}
	log = log.With(zap.String("program", programName))

	log.Info("init Zap logger success!")

	return log, nil
}
