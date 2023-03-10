package handler

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
)

type Option func(*handler) error

func WithService(svc newsletter.Service) Option {
	return func(h *handler) error {
		h.svc = svc
		return nil
	}
}
