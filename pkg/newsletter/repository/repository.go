package repository

import (
	"sync"

	newsletter "github.com/Eitol/newsletter-api/pkg/newsletter"
)

var (
	repo newsletter.Repository
	once sync.Once
)

type repository struct {
	data []*subscriptionDBModel
}

type Option func(*repository) error

func New(opts ...Option) (newsletter.Repository, error) {
	r := &repository{}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func Must() newsletter.Repository {
	once.Do(func() {
		r, err := New()
		if err != nil {
			panic(err)
		}

		repo = r
	})

	return repo
}
