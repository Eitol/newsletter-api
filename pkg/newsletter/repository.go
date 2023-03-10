package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Search(
		ctx context.Context,
		userID uuid.UUID,
		blogID uuid.UUID,
		interests []Interest,
		limit int,
		offset int,
	) ([]*Subscription, error)
}
