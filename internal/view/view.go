package view

import (
	"l0/internal/domain"

	"go.uber.org/zap"
)

type View struct {
	Log    *zap.Logger
	Domain *domain.Domain
}

func NewView(log *zap.Logger, domain *domain.Domain) *View {
	return &View{
		Log:    log,
		Domain: domain,
	}
}

type ViewIFace interface {
}
