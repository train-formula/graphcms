package mutation

import "context"

func (r *MutationResolver) Health(ctx context.Context) (string, error) {
	return "OK", nil
}
