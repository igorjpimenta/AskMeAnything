package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/igorjpimenta/AskMeAnything/internal/store/pgstore"
	"github.com/jackc/pgx/v5"
)

func (h apiHandler) handleError(w http.ResponseWriter, msg string, err error, status int) {
	slog.Warn(msg, "error", err)
	if status == http.StatusInternalServerError {
		msg = "something went wrong"
	}
	http.Error(w, msg, status)
}

func (h apiHandler) decodeJSONBody(w http.ResponseWriter, r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		h.handleError(w, "invalid json", err, http.StatusBadRequest)
		return err
	}
	return nil
}

func (h apiHandler) respondWithJSON(w http.ResponseWriter, status int, response any) {
	data, err := json.Marshal(response)
	if err != nil {
		h.handleError(w, "failed to encode response", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		h.handleError(w, "failed to write the response", err, http.StatusInternalServerError)
		return
	}
}

func getEntityByID[T any](h apiHandler, w http.ResponseWriter, r *http.Request, param string, fetchFunc func(ctx context.Context, id uuid.UUID) (entity T, err error)) (entity T, id uuid.UUID, ok bool) {
	var zero T
	rawID := chi.URLParam(r, param)
	id, err := uuid.Parse(rawID)
	if err != nil {
		h.handleError(w, "invalid "+param, err, http.StatusBadRequest)
		return zero, uuid.Nil, false
	}

	entity, err = fetchFunc(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			h.handleError(w, param+" not found", err, http.StatusBadRequest)
			return zero, uuid.Nil, false
		}
		h.handleError(w, "something went wrong", err, http.StatusInternalServerError)
		return zero, uuid.Nil, false
	}

	return entity, id, true
}

func (h apiHandler) getRoom(w http.ResponseWriter, r *http.Request) (room pgstore.Room, roomID uuid.UUID, ok bool) {
	return getEntityByID(h, w, r, "room_id", h.q.GetRoom)
}

func (h apiHandler) getMessage(w http.ResponseWriter, r *http.Request) (message pgstore.Message, messageID uuid.UUID, ok bool) {
	return getEntityByID(h, w, r, "message_id", h.q.GetMessage)
}

func (h apiHandler) validateMessageRoom(w http.ResponseWriter, r *http.Request) (message pgstore.Message, roomID uuid.UUID, ok bool) {
	message, _, ok = h.getMessage(w, r)
	if !ok {
		return
	}
	_, roomID, ok = h.getRoom(w, r)
	if !ok {
		return
	}

	if message.RoomID != roomID {
		h.handleError(w, "message does not belong to the specified room", nil, http.StatusNotFound)
		return pgstore.Message{}, uuid.Nil, false
	}

	return message, roomID, true
}
