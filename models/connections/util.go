package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
)

type TotalCountResolver func(ctx context.Context) (int, error)

func GeneratePageInfo(numEdges int, cursorRetriever func() (startCursor, endCursor string, err error)) (*models.PageInfo, error) {

	if numEdges > 0 {
		startCursor, endCursor, err := cursorRetriever()
		if err != nil {
			return nil, err
		}

		return &models.PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
		}, nil
	}

	return &models.PageInfo{}, nil
}
