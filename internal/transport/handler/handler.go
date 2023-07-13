package handler

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
	MessageService
}

func NewHandler(log *log.Logger, m MessageService) *Handler {
	return &Handler{
		logger:         log,
		MessageService: m,
	}
}

func (h Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/getMessage", h.getMessage)

	return mux
}
