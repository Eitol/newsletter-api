package service

import (
	"context"
	"github.com/Eitol/newsletter-api/pkg/newsletter"

	"github.com/google/uuid"
)

func (s *service) Post(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
) (uuid.UUID, error) {
	r, err := s.repo.Create(ctx, userID, blogID, interests)
	if err != nil {
		return uuid.Nil, err
	}
	return r, nil
}
