package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"text-cdn/internal/core"
)

type MessageService interface {
	GetMessage(ctx context.Context, msg core.Message) (core.Message, error)
}

func (h Handler) getMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(h.responseWithError(http.StatusMethodNotAllowed, MethodNotAllowed))
		return
	}

	var in core.Message
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.Write(h.responseWithError(http.StatusBadRequest, BadJSON))
		return
	}

	if err := h.validateMessage(in); err != nil {
		w.Write(h.responseWithError(http.StatusBadRequest, BadJSON))
		return
	}

	msg, err := h.MessageService.GetMessage(context.TODO(), in)
	if err != nil {
		w.Write(h.responseWithError(http.StatusNotFound, NotFound))
		return
	}

	responseJSON, err := json.Marshal(&msg)
	if err != nil {
		// log
		return
	}
	w.Write(responseJSON)
}

func (h Handler) validateMessage(in core.Message) error {
	if in.Type == "" || in.Language == "" {
		return errors.New(SomeJSONFiledsEmpty)
	}
	return nil
}
