package repository

import (
	"context"

	"github.com/Eitol/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	_ context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) (*newsletter.SearchResult, error) {
	mutex.Lock()
	defer mutex.Unlock()
	resultIdxPool = resultIdxPool[:0]
	for i, v := range inMemoryDB {
		matchUserUUID := true
		if userID != uuid.Nil && v.UserID != userID {
			matchUserUUID = false
		}
		if !matchUserUUID {
			continue
		}
		matchBlogUUID := true
		if blogID != uuid.Nil && v.BlogID != blogID {
			matchBlogUUID = false
		}
		if !matchBlogUUID {
			continue
		}
		matchAllInterests := true
		for _, interest := range interests {
			found := false
			for _, vInterest := range v.Interests {
				if interest == vInterest {
					found = true
					break
				}
			}
			if !found {
				matchAllInterests = false
				break
			}
		}
		if !matchAllInterests {
			continue
		}
		resultIdxPool = append(resultIdxPool, i)
	}
	total := len(resultIdxPool)
	pages := total / limit // integer division
	if total%limit != 0 {
		pages++
	}
	searchResult := &newsletter.SearchResult{
		Subscriptions: nil,
		Total:         total,
		Pages:         pages,
	}
	if offset >= len(resultIdxPool) {
		return searchResult, nil
	}
	startIdx := offset
	endIdx := offset + limit
	if endIdx > len(resultIdxPool) {
		endIdx = len(resultIdxPool)
	}
	result := make([]*newsletter.Subscription, endIdx-startIdx)
	for i, v := range resultIdxPool[startIdx:endIdx] {
		result[i] = &inMemoryDB[v]
	}
	searchResult.Subscriptions = result
	return searchResult, nil
}
