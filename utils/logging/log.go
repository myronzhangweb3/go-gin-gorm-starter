package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel            = viper.GetString("LOG.LEVEL")
	defaultLogLevelInfo = "info"

	Log *zap.Logger
)

func InitLogger(programName string) {
	InitLoggerWithLogFile(programName, "")
}

func InitLoggerWithLogFile(programName, logFilePath string) {
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
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.Encoding = "json"

	// Create a file to write logs
	if logFilePath != "" {
		// Ensure the directory exists
		dir := filepath.Dir(logFilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("%s init_db logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return
		}

		// Create a file to write logs
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("%s init_db logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return
		}
		// Create a core that writes logs to the file
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(file),
			level,
		)
		// Build the logger with the file core
		Log = zap.New(fileCore, zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		var err error
		Log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			fmt.Printf("%s init_db logger error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
			return
		}
	}
	Log = Log.With(zap.String("program", programName))

	Log.Info("init Zap logger success!")
}
