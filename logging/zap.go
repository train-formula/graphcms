package logging

import (
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
	"go.uber.org/zap"
)

func UUID(key string, uuid uuid.UUID) zap.Field {
	return zap.String(key, uuid.String())
}

func Cursor(key string, cursor cursor.Cursor) zap.Field {
	return zap.String(key, cursor.Serialize())
}
