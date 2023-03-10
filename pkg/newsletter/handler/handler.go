package handler

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
)

type handler struct {
	svc newsletter.Service
}

func New(opts ...Option) (newsletter.Handler, error) {
	h := &handler{}

	for _, opt := range opts {
		if err := opt(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

func Must(
	svc newsletter.Service,
) newsletter.Handler {
	h, err := New(WithService(svc))
	if err != nil {
		panic(err)
	}

	return h
}
