package internal

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	Logger = zap.NewExample()
	defer Logger.Sync()
}
