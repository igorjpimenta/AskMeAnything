package api

import (
	"net/http"

	"github.com/igorjpimenta/AskMeAnything/internal/store/pgstore"
)

func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Theme string `json:"theme"`
	}
	if err := h.decodeJSONBody(w, r, &body); err != nil {
		return
	}

	roomID, err := h.q.InsertRoom(r.Context(), body.Theme)
	if err != nil {
		h.handleError(w, "failed to insert room", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		ID         string `json:"id"`
		OwnerToken string `json:"owner_token"`
	}

	h.respondWithJSON(w, http.StatusOK, response{ID: roomID.String()})
}

func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.q.GetRooms(r.Context())
	if err != nil {
		h.handleError(w, "failed to get rooms", err, http.StatusInternalServerError)
		return
	}

	if rooms == nil {
		rooms = []pgstore.Room{}
	}

	h.respondWithJSON(w, http.StatusOK, rooms)
}

func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {
	_, roomID, ok := h.getRoom(w, r)
	if !ok {
		return
	}

	var body struct {
		Message string `json:"message"`
	}
	if err := h.decodeJSONBody(w, r, &body); err != nil {
		return
	}

	messageID, err := h.q.InsertMessage(r.Context(), pgstore.InsertMessageParams{RoomID: roomID, Message: body.Message})
	if err != nil {
		h.handleError(w, "failed to insert message", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	h.respondWithJSON(w, http.StatusOK, response{ID: messageID.String()})

	go h.notifyClients(Message{
		Kind:   MessageKindMessageCreated,
		RoomID: roomID.String(),
		Value: MessageMessageCreated{
			ID:      messageID.String(),
			Message: body.Message,
		},
	})
}

func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {
	_, roomID, ok := h.getRoom(w, r)
	if !ok {
		return
	}

	messages, err := h.q.GetRoomMessages(r.Context(), roomID)
	if err != nil {
		h.handleError(w, "failed to get messages from the room", err, http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []pgstore.Message{}
	}

	h.respondWithJSON(w, http.StatusOK, messages)
}

func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request) {
	message, _, ok := h.validateMessageRoom(w, r)
	if !ok {
		return
	}

	h.respondWithJSON(w, http.StatusOK, message)
}

func (h apiHandler) handleReactToMessage(w http.ResponseWriter, r *http.Request) {
	message, roomID, ok := h.validateMessageRoom(w, r)
	if !ok {
		return
	}

	reactionCount, err := h.q.ReactToMessage(r.Context(), message.ID)
	if err != nil {
		h.handleError(w, "failed to react to message", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		Count int64 `json:"count"`
	}

	h.respondWithJSON(w, http.StatusOK, response{Count: reactionCount})

	go h.notifyClients(Message{
		Kind:   MessageKindMessageReactionIncreased,
		RoomID: roomID.String(),
		Value: MessageMessageReactionIncreased{
			ID:    message.ID.String(),
			Count: reactionCount,
		},
	})
}

func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request) {
	message, roomID, ok := h.validateMessageRoom(w, r)
	if !ok {
		return
	}

	reactionCount, err := h.q.RemoveReactionFromMessage(r.Context(), message.ID)
	if err != nil {
		h.handleError(w, "failed to react to message", err, http.StatusInternalServerError)
		return
	}

	type response struct {
		Count int64 `json:"count"`
	}

	h.respondWithJSON(w, http.StatusOK, response{Count: reactionCount})

	go h.notifyClients(Message{
		Kind:   MessageKindMessageReactionDecreased,
		RoomID: roomID.String(),
		Value: MessageMessageReactionDecreased{
			ID:    message.ID.String(),
			Count: reactionCount,
		},
	})
}

func (h apiHandler) handleMaskMessageAsAnswered(w http.ResponseWriter, r *http.Request) {
	message, roomID, ok := h.validateMessageRoom(w, r)
	if !ok {
		return
	}

	if message.Answered {
		h.handleError(w, "message already answered", nil, http.StatusBadRequest)
		return
	}

	err := h.q.MarkMessageAnswered(r.Context(), message.ID)
	if err != nil {
		h.handleError(w, "failed to mark message answered", err, http.StatusInternalServerError)
		return
	}

	go h.notifyClients(Message{
		Kind:   MessageKindMessageAnswered,
		RoomID: roomID.String(),
		Value: MessageMessageAnswered{
			ID: message.ID.String(),
		},
	})
}
