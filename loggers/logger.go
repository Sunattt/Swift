package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (*zap.Logger, error) {
	//экземпляр конфигурации логера
	config := zap.NewProductionConfig()
	//установка времени форматирования по стондарту ISO8601
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//устанавливается путь, куда будут выводиться логи
	config.OutputPaths = []string{"./loggers/logger.log"}
	//постороение экземпляра логгера
	return config.Build()
}
