package repository

import (
	"context"
	"github.com/Eitol/newsletter-api/pkg/newsletter"

	"github.com/google/uuid"
)

func (r *repository) Create(
	_ context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
) (uuid.UUID, error) {
	id := uuid.New()
	inMemoryDB = append(inMemoryDB, newsletter.Subscription{
		UserID:    userID,
		BlogID:    blogID,
		Interests: interests,
	})
	return id, nil
}
