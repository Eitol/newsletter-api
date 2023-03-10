//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"github.com/Eitol/newsletter-api/pkg/newsletter/repository"
	"github.com/Eitol/newsletter-api/pkg/newsletter/service"
	"github.com/google/wire"
)

func Build() newsletter.Handler {
	wire.Build(
		Must,
		service.Must,
		repository.Must,
	)

	return nil
}
