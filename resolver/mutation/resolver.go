package mutation

import "github.com/go-pg/pg/v9"

func NewMutationResolver(db *pg.DB) *MutationResolver {
	return &MutationResolver{
		db: db,
	}
}

type MutationResolver struct {
	db *pg.DB
}
