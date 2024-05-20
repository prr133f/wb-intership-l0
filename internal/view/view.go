package view

import (
	"l0/internal/domain"

	"go.uber.org/zap"
)

type View struct {
	Log    *zap.Logger
	Domain domain.IFace
}

func NewView(log *zap.Logger, domain domain.IFace) *View {
	return &View{
		Log:    log,
		Domain: domain,
	}
}
