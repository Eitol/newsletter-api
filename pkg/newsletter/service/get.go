package service

import (
	"context"

	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	page int,
	maxPageSize int,
) (*newsletter.Result[*newsletter.Subscription], error) {
	limit := maxPageSize
	offset := (page - 1) * maxPageSize
	r, err := s.repo.Search(ctx, userID, blogID, interests, limit, offset)
	if err != nil {
		return nil, err
	}
	return &newsletter.Result[*newsletter.Subscription]{
		Total: r.Total,
		Pages: r.Pages,
		Page: newsletter.Page[*newsletter.Subscription]{
			Number:   page,
			Elements: r.Subscriptions,
		},
	}, nil
}
