package handler

import (
	"log/slog"

	"github.com/su-its/typing/typing-server/internal/infra/ent"
)

type Handler struct {
	log       *slog.Logger
	entClient *ent.Client
}

func New(log *slog.Logger, entClient *ent.Client) *Handler {
	return &Handler{
		log:       log,
		entClient: entClient,
	}
}
