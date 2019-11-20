package logging

import "go.uber.org/zap"

func LogRollbackError(logger *zap.Logger, err error) {

	if err != nil {
		logger.Error("Failed to rollback transaction", zap.Error(err))
	}
}
