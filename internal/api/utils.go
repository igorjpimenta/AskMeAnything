package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

func (h apiHandler) getRoom(w http.ResponseWriter, r *http.Request) (roomID uuid.UUID, ok bool) {
	rawRoomID := chi.URLParam(r, "room_id")
	roomID, err := uuid.Parse(rawRoomID)
	if err != nil {
		h.handleError(w, "invalid room id", err, http.StatusBadRequest)
		return uuid.Nil, false
	}

	_, err = h.q.GetRoom(r.Context(), roomID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			h.handleError(w, "room not found", err, http.StatusBadRequest)
			return uuid.Nil, false
		}

		h.handleError(w, "something went wrong", err, http.StatusInternalServerError)
		return uuid.Nil, false
	}

	return roomID, true
}

func (h apiHandler) notifyClients(msg Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	subscribers, ok := h.subscribers[msg.RoomID]
	if !ok || len(subscribers) == 0 {
		return
	}

	for conn, cancel := range subscribers {
		if err := conn.WriteJSON(msg); err != nil {
			slog.Error("failed to send message to client", "error", err)
			cancel()
		}
	}
}
