package logging

import (
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

func UUID(key string, uuid uuid.UUID) zap.Field {
	return zap.String(key, uuid.String())
}
