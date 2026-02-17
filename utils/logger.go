package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
)

var logger *logs.BeeLogger

func InitLogger(logPath string) error {
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return err
	}

	logger = logs.NewLogger(10000)
	logger.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","maxsize":10485760,"maxdays":30,"level":6}`, logPath))
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)

	return nil
}

func GetLogger() *logs.BeeLogger {
	if logger == nil {
		logger = logs.NewLogger(10000)
		logger.SetLogger(logs.AdapterConsole)
	}
	return logger
}

func LogInfo(format string, v ...interface{}) {
	GetLogger().Info(format, v...)
}

func LogWarn(format string, v ...interface{}) {
	GetLogger().Warning(format, v...)
}

func LogError(format string, v ...interface{}) {
	GetLogger().Error(format, v...)
}

func LogAccess(ip, method, path, userAgent string, statusCode int, responseTime int64) {
	GetLogger().Info("Access | %s | %s | %s | %s | %d | %dms",
		ip, method, path, userAgent, statusCode, responseTime)
}
