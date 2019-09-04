package tagdb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/tag"
)

func InsertTag(ctx context.Context, conn database.Conn, new tag.Tag) (*tag.Tag, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}
