package query

import "context"

func (r *QueryResolver) Health(ctx context.Context) (string, error) {
	return "OK", nil
}
