package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Get(
		ctx context.Context,
		UserID uuid.UUID,
		BlogID uuid.UUID,
		Interests []Interest,
	) (*Result[*Subscription], error)

	Post(
		ctx context.Context,
		userID uuid.UUID,
		blogID uuid.UUID,
		interests []Interest,
	) (uuid.UUID, error)
}
