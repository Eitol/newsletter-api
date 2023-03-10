package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type SearchResult struct {
	Subscriptions []*Subscription
	Total         int
	Pages         int
}

type Repository interface {
	Search(
		ctx context.Context,
		userID uuid.UUID,
		blogID uuid.UUID,
		interests []Interest,
		limit int,
		offset int,
	) (*SearchResult, error)

	Create(
		ctx context.Context,
		userID uuid.UUID,
		blogID uuid.UUID,
		interests []Interest,
	) (uuid.UUID, error)
}
