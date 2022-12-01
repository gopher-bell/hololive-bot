package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLog *zap.SugaredLogger

func SetupZap() (func(), error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	var err error
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	ZapLog = logger.Sugar()

	return func() {
		logger.Sync()
	}, nil
}
