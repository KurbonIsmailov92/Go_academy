package logger

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
	Debug   *log.Logger
)

const (
	LogInfo       = "logs/info.log"
	LogError      = "logs/error.log"
	LogWarning    = "logs/warning.log"
	LogDebug      = "logs/debug.log"
	LogMaxSize    = 25
	LogMaxBackups = 5
	LogMaxAge     = 30
	LogCompress   = true
)

func Init() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return err
		}
	}

	lumberLogInfo := &lumberjack.Logger{
		Filename:   LogInfo,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   LogCompress,
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   LogError,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   LogCompress,
		LocalTime:  true,
	}

	lumberLogWarning := &lumberjack.Logger{
		Filename:   LogWarning,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   LogCompress,
		LocalTime:  true,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   LogDebug,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   LogCompress,
		LocalTime:  true,
	}

	Info = log.New(lumberLogInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(lumberLogError, "", log.Ldate|log.Lmicroseconds)
	Warning = log.New(lumberLogWarning, "", log.Ldate|log.Lmicroseconds)
	Debug = log.New(lumberLogDebug, "", log.Ldate|log.Lmicroseconds)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	Info.SetOutput(gin.DefaultWriter)
	Error.SetOutput(lumberLogError)
	Warning.SetOutput(lumberLogWarning)
	Debug.SetOutput(lumberLogDebug)

	return nil
}
