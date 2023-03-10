package service

import (
	"context"

	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	UserID uuid.UUID,
	BlogID uuid.UUID,
	Interests []newsletter.Interest,
) (*newsletter.Result[*newsletter.Subscription], error) {
	panic("implement me")
}
