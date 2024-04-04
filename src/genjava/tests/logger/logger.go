package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func InitLogger(logWriter io.Writer) *zap.Logger {
	writeSyncer := zapcore.AddSync(logWriter)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writeSyncer,
		zap.InfoLevel,
	)
	return zap.New(core)
}
